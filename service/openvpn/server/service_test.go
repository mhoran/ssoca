package server_test

import (
	"net/http"

	"github.com/dpb587/ssoca/server/service"
	. "github.com/dpb587/ssoca/service/openvpn/server"
	svcconfig "github.com/dpb587/ssoca/service/openvpn/server/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service", func() {
	var subject Service

	Describe("interface", func() {
		It("github.com/dpb587/ssoca/server/service.Service", func() {
			var _ service.Service = (*Service)(nil)
		})
	})

	Context("Basics", func() {
		BeforeEach(func() {
			subject = NewService("test1", svcconfig.Config{})
		})

		Describe("Name", func() {
			It("works", func() {
				Expect(subject.Name()).To(Equal("test1"))
			})
		})

		Describe("Metadata", func() {
			It("works", func() {
				Expect(subject.Metadata()).To(BeNil())
			})
		})

		Describe("VerifyAuthorization", func() {
			It("works", func() {
				err := subject.VerifyAuthorization(http.Request{}, nil)
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})
})
