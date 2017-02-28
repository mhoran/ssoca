package config_test

import (
	. "github.com/dpb587/ssoca/service/ssh/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	var subject Config

	BeforeEach(func() {
		subject = Config{}
	})

	Describe("ApplyDefaults", func() {
		Describe("with defaults", func() {
			BeforeEach(func() {
				subject.ApplyDefaults()
			})

			It("defaults certauth: default", func() {
				Expect(subject.CertAuthName).To(Equal("default"))
			})

			It("defaults validity", func() {
				Expect(subject.ValidityString).To(Equal("2m"))
			})

			It("defaults extensions", func() {
				Expect(subject.Extensions).To(Equal(ExtensionDefaults))
			})
		})

		Describe("not really dealing with defaults but its convenient so...", func() {
			It("ignores extensions if ssoca-no-defaults is used", func() {
				subject.Extensions = Extensions{ExtensionNoDefaults}
				subject.ApplyDefaults()

				Expect(subject.Extensions).To(HaveLen(0))
			})
		})
	})

	Describe("Target", func() {
		Describe("Configured", func() {
			It("defaults to not", func() {
				Expect(subject.Target.Configured()).To(BeFalse())
			})

			Context("a property is configured", func() {
				Context("host", func() {
					It("is configured", func() {
						subject.Target.Host = "localhost"

						Expect(subject.Target.Configured()).To(BeTrue())
					})
				})

				Context("user", func() {
					It("is configured", func() {
						subject.Target.RawUser = "vcap"

						Expect(subject.Target.Configured()).To(BeTrue())
					})
				})

				Context("port", func() {
					It("is configured", func() {
						subject.Target.Port = 2222

						Expect(subject.Target.Configured()).To(BeTrue())
					})
				})
			})
		})
	})
})
