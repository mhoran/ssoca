package or_test

import (
	"errors"
	"net/http"

	"github.com/dpb587/ssoca/auth"
	. "github.com/dpb587/ssoca/authz/filter/or"

	"github.com/dpb587/ssoca/authz/filter"
	"github.com/dpb587/ssoca/authz/filter/filterfakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Requirement", func() {
	var request http.Request
	var token *auth.Token
	var subject Requirement

	satisfyRequirement := &filterfakes.FakeRequirement{}
	satisfyRequirement.IsSatisfiedReturns(true, nil)

	dissatisfyRequirement := &filterfakes.FakeRequirement{}
	dissatisfyRequirement.IsSatisfiedReturns(false, nil)

	errorRequirement := &filterfakes.FakeRequirement{}
	errorRequirement.IsSatisfiedReturns(false, errors.New("faked error"))

	BeforeEach(func() {
		request = http.Request{}
		token = &auth.Token{ID: "test"}
	})

	Describe("IsSatisfied", func() {
		Context("without filters", func() {
			BeforeEach(func() {
				subject = Requirement{
					Requirements: []filter.Requirement{},
				}
			})

			It("does not satisfy", func() {
				satisfied, err := subject.IsSatisfied(&request, token)

				Expect(err).ToNot(HaveOccurred())
				Expect(satisfied).To(BeFalse())
			})
		})

		Context("with satisfying filters", func() {
			BeforeEach(func() {
				subject = Requirement{
					Requirements: []filter.Requirement{
						satisfyRequirement,
						satisfyRequirement,
					},
				}
			})

			It("satisfies", func() {
				satisfied, err := subject.IsSatisfied(&request, token)

				Expect(err).ToNot(HaveOccurred())
				Expect(satisfied).To(BeTrue())
			})
		})

		Context("with unsatisfying filters", func() {
			BeforeEach(func() {
				subject = Requirement{
					Requirements: []filter.Requirement{
						dissatisfyRequirement,
					},
				}
			})

			It("does not satisfy", func() {
				satisfied, err := subject.IsSatisfied(&request, token)

				Expect(err).ToNot(HaveOccurred())
				Expect(satisfied).To(BeFalse())
			})
		})

		Context("with some dissatisfying filters", func() {
			It("satisfies", func() {
				subject = Requirement{
					Requirements: []filter.Requirement{
						satisfyRequirement,
						dissatisfyRequirement,
					},
				}

				satisfied, err := subject.IsSatisfied(&request, token)

				Expect(err).ToNot(HaveOccurred())
				Expect(satisfied).To(BeTrue())
			})

			It("stops evaluating early", func() {
				subject = Requirement{
					Requirements: []filter.Requirement{
						satisfyRequirement,
						dissatisfyRequirement,
						errorRequirement,
					},
				}

				satisfied, err := subject.IsSatisfied(&request, token)

				Expect(err).ToNot(HaveOccurred())
				Expect(satisfied).To(BeTrue())
			})
		})

		Context("with some erroring filters", func() {
			BeforeEach(func() {
				subject = Requirement{
					Requirements: []filter.Requirement{
						errorRequirement,
						satisfyRequirement,
					},
				}
			})

			It("does not satisfy", func() {
				satisfied, err := subject.IsSatisfied(&request, token)

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("faked error"))
				Expect(satisfied).To(BeFalse())
			})
		})
	})
})
