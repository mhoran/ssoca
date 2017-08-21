package fs_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"math/big"

	"github.com/sirupsen/logrus"
	logrustest "github.com/sirupsen/logrus/hooks/test"
	"golang.org/x/crypto/ssh"

	boshsysfakes "github.com/cloudfoundry/bosh-utils/system/fakes"
	. "github.com/dpb587/ssoca/certauth/fs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Provider", func() {
	var subject Provider
	var fs boshsysfakes.FakeFileSystem
	var logger logrus.FieldLogger

	// certstrap init --key-bits 1024 --common-name ssoca-test --passphrase ''
	var ca1crtStr = `-----BEGIN CERTIFICATE-----
MIIB5TCCAU6gAwIBAgIBATANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwpzc29j
YS10ZXN0MB4XDTE3MDIxMzIwMzMwOFoXDTI3MDIxMzIwMzMwOFowFTETMBEGA1UE
AxMKc3NvY2EtdGVzdDCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA6Td3vsA/
f9lHQPeIzJB9J1JTXaDVfqoClU2ZRlua7BmlfXOQngo/1OmetO3THEr+mxFMGgfJ
Z6CujBRd3A7/3h+Iw72jKGBag4iEBI9uBcyeRgWdzcV7l7dzQT00XCBnkJJRJV4j
oDIovquAz6iKf4Al4wyQ5k1RM9KAlUipIFkCAwEAAaNFMEMwDgYDVR0PAQH/BAQD
AgEGMBIGA1UdEwEB/wQIMAYBAf8CAQAwHQYDVR0OBBYEFP8lIbNl3zZPEHF17cFU
NFsK/0/oMA0GCSqGSIb3DQEBCwUAA4GBADMCd4nzc19voa60lNknhsihcfyNUeUt
EEsLCceK+9F1u2Xdj+mTNOh3MI+5m7wmFLiHuUtQovHMJ4xUpoHa6Iznc+QCbow4
SMO3sf1847tASv3eUFwEUt9vv39vtey6C6ftiUUImzZYfx6FO/A62uGEg2w3IOJ+
3cCXYiulfsyv
-----END CERTIFICATE-----`
	var ca1keyStr = `-----BEGIN RSA PRIVATE KEY-----
MIICXwIBAAKBgQDpN3e+wD9/2UdA94jMkH0nUlNdoNV+qgKVTZlGW5rsGaV9c5Ce
Cj/U6Z607dMcSv6bEUwaB8lnoK6MFF3cDv/eH4jDvaMoYFqDiIQEj24FzJ5GBZ3N
xXuXt3NBPTRcIGeQklElXiOgMii+q4DPqIp/gCXjDJDmTVEz0oCVSKkgWQIDAQAB
AoGBANC3T3drXmjw74/4+Hj7Jsa2Kt20Pt1pEX7FP9Nz0CZUnYK0lkyaJ55IpjyO
S00a4NmulUkGhv0zFINRBt8WnW1bjBxNmqyBYh2diO3vA/gk8U1gcifW1LQt8WmE
ietvN3OFXI1a7FipchCZYQn5Rr8O3a/tjwohtWIDdaDltw+xAkEA7Ybxu8OXQnvy
Y+fDISRGG5vDFGnNGe9KcREIxSF6LWJ7+ap5LmMxnhfag5qlrObQW3K2miTpGYkl
CIRRNFMIvwJBAPtatE1evu25R3NSTU2YwQgkEymh40PW+lncYge6ZqZGfK7J5JBK
wr1ug7KjTJgIfY2Sg2VHn56HAdA4RUl2xOcCQQDZqnTxpQ6DHYSFqwg04cHhYP8H
QOF0Z8WnEX4g8Em/N2X26BK+wKXig2d6fIhghu/fLaNKZJK8FOK8CE1GDuWPAkEA
wrP6Ysx3vZH+JPil5Ovk6zd2mJNMhmpqt10dmrrrdPW483R01sjynOaUobYZSNOa
3iWWHsgifxw5bV+JXGTiFQJBAKwh6Hvli5hcfoepPMz2RQnmU1NM8hJOHHeZh+eT
z6hlMpOS9rSjABcBdXxXjFXtIEjWUG5Tj8yOYd735zY8Ny8=
-----END RSA PRIVATE KEY-----`

	pemToCertificate := func(bytes []byte) x509.Certificate {
		pem, _ := pem.Decode(bytes)
		if pem == nil {
			panic("Failed decoding PEM")
		}

		certificate, err := x509.ParseCertificate(pem.Bytes)
		if err != nil {
			panic(err)
		}

		return *certificate
	}

	BeforeEach(func() {
		fs = *boshsysfakes.NewFakeFileSystem()
		fs.WriteFileString("/path/ca1/certificate", ca1crtStr)
		fs.WriteFileString("/path/ca1/private_key", ca1keyStr)
		fs.WriteFileString("/path/ca2/certificate", "broken")
		fs.WriteFileString("/path/ca2/private_key", "broken")

		logger, _ = logrustest.NewNullLogger()
	})

	Describe("SignCertificate", func() {
		var testKey *rsa.PrivateKey
		var template x509.Certificate

		BeforeEach(func() {
			var err error

			testKey, err = rsa.GenerateKey(rand.Reader, 1024)
			if err != nil {
				Fail("generating private key")
			}

			template = x509.Certificate{
				SerialNumber: big.NewInt(12345),
				Subject: pkix.Name{
					CommonName: "ssoca-fake1",
				},
			}
		})

		It("signs certificate", func() {
			subject = NewProvider(
				"name1",
				Config{
					CertificatePath: "/path/ca1/certificate",
					PrivateKeyPath:  "/path/ca1/private_key",
				},
				&fs,
				logger,
			)

			bytes, err := subject.SignCertificate(&template, &testKey.PublicKey, logrus.Fields{})

			Expect(err).ToNot(HaveOccurred())
			Expect(len(bytes)).To(BeNumerically(">", 0))

			certificate := pemToCertificate(bytes)
			Expect(certificate.SerialNumber).To(BeEquivalentTo(big.NewInt(12345)))
			Expect(certificate.Subject.CommonName).To(Equal("ssoca-fake1"))

			caCertificate := pemToCertificate([]byte(ca1crtStr))
			err = certificate.CheckSignatureFrom(&caCertificate)

			Expect(err).ToNot(HaveOccurred())
		})

		Context("certificate/key errors", func() {
			It("errors on missing certificate", func() {
				subject = NewProvider(
					"name1",
					Config{
						CertificatePath: "/path/ca0/certificate",
						PrivateKeyPath:  "/path/ca1/private_key",
					},
					&fs,
					logger,
				)

				_, err := subject.SignCertificate(&template, &testKey.PublicKey, logrus.Fields{})

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Getting CA certificate"))
			})

			It("errors on missing private key", func() {
				subject = NewProvider(
					"name1",
					Config{
						CertificatePath: "/path/ca1/certificate",
						PrivateKeyPath:  "/path/ca0/private_key",
					},
					&fs,
					logger,
				)

				_, err := subject.SignCertificate(&template, &testKey.PublicKey, logrus.Fields{})

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Getting CA private key"))
			})

			It("errors on misconfigured private key", func() {
				subject = NewProvider(
					"name1",
					Config{
						CertificatePath: "/path/ca1/certificate",
						PrivateKeyPath:  "/path/ca1/certificate",
					},
					&fs,
					logger,
				)

				_, err := subject.SignCertificate(&template, &testKey.PublicKey, logrus.Fields{})

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Getting CA private key"))
			})

			It("errors on invalid private key", func() {
				subject = NewProvider(
					"name1",
					Config{
						CertificatePath: "/path/ca1/certificate",
						PrivateKeyPath:  "/path/ca2/certificate",
					},
					&fs,
					logger,
				)

				_, err := subject.SignCertificate(&template, &testKey.PublicKey, logrus.Fields{})

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Getting CA private key"))
			})
		})
	})

	Describe("SignSSHCertificate", func() {
		var testKey *rsa.PrivateKey
		var cert ssh.Certificate

		BeforeEach(func() {
			var err error

			testKey, err = rsa.GenerateKey(rand.Reader, 1024)
			if err != nil {
				Fail("generating private key")
			}

			publicKey, err := ssh.NewPublicKey(&testKey.PublicKey)
			if err != nil {
				Fail("parsing to public key")
			}

			cert = ssh.Certificate{
				Nonce:    []byte("ssoca-fake1"),
				Key:      publicKey,
				CertType: ssh.UserCert,
			}
		})

		It("signs certificate", func() {
			subject = NewProvider(
				"name1",
				Config{
					CertificatePath: "/path/ca1/certificate",
					PrivateKeyPath:  "/path/ca1/private_key",
				},
				&fs,
				logger,
			)

			Expect(cert.Signature).To(BeNil())

			err := subject.SignSSHCertificate(&cert, logrus.Fields{})
			Expect(err).ToNot(HaveOccurred())

			// @todo use Verify instead
			Expect(cert.Signature).ToNot(BeNil())
		})

		Context("certificate/key errors", func() {
			It("errors on missing private key", func() {
				subject = NewProvider(
					"name1",
					Config{
						CertificatePath: "/path/ca1/certificate",
						PrivateKeyPath:  "/path/ca0/private_key",
					},
					&fs,
					logger,
				)

				err := subject.SignSSHCertificate(&cert, logrus.Fields{})

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Getting CA private key"))
			})

			It("errors on misconfigured private key", func() {
				subject = NewProvider(
					"name1",
					Config{
						CertificatePath: "/path/ca1/certificate",
						PrivateKeyPath:  "/path/ca1/certificate",
					},
					&fs,
					logger,
				)

				err := subject.SignSSHCertificate(&cert, logrus.Fields{})

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Getting CA private key"))
			})

			It("errors on invalid private key", func() {
				subject = NewProvider(
					"name1",
					Config{
						CertificatePath: "/path/ca1/certificate",
						PrivateKeyPath:  "/path/ca2/certificate",
					},
					&fs,
					logger,
				)

				err := subject.SignSSHCertificate(&cert, logrus.Fields{})

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Getting CA private key"))
			})
		})
	})

	Describe("GetCertificate", func() {
		It("provides certificate", func() {
			subject = NewProvider(
				"name1",
				Config{
					CertificatePath: "/path/ca1/certificate",
					PrivateKeyPath:  "/path/ca1/private_key",
				},
				&fs,
				logger,
			)

			crt, err := subject.GetCertificate()

			Expect(err).ToNot(HaveOccurred())
			Expect(crt).To(BeAssignableToTypeOf(&x509.Certificate{}))
			Expect(crt.IsCA).To(BeTrue())
			Expect(crt.Subject.CommonName).To(Equal("ssoca-test"))
		})

		Context("filesystem errors", func() {
			It("errors", func() {
				subject = NewProvider(
					"name1",
					Config{
						CertificatePath: "/path/ca0/certificate",
						PrivateKeyPath:  "/path/ca0/private_key",
					},
					&fs,
					logger,
				)

				_, err := subject.GetCertificate()

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Reading certificate"))
			})
		})

		Context("bad certificate contents", func() {
			It("errors", func() {
				subject = NewProvider(
					"name1",
					Config{
						CertificatePath: "/path/ca2/certificate",
						PrivateKeyPath:  "/path/ca2/private_key",
					},
					&fs,
					logger,
				)

				_, err := subject.GetCertificate()

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Failed decoding certificate PEM"))
			})
		})

		Context("invalid certificate format", func() {
			It("errors", func() {
				subject = NewProvider(
					"name1",
					Config{
						CertificatePath: "/path/ca1/private_key",
						PrivateKeyPath:  "/path/ca1/private_key",
					},
					&fs,
					logger,
				)

				_, err := subject.GetCertificate()

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Parsing certificate"))
			})
		})
	})

	Describe("GetCertificatePEM", func() {
		var subject Provider

		BeforeEach(func() {
			subject = NewProvider(
				"name1",
				Config{
					CertificatePath: "/path/ca1/certificate",
					PrivateKeyPath:  "/path/ca1/private_key",
				},
				&fs,
				logger,
			)
		})

		It("provides string", func() {
			crt, err := subject.GetCertificatePEM()

			Expect(err).ToNot(HaveOccurred())
			Expect(crt).To(Equal(ca1crtStr))
		})

		Context("with filesystem error", func() {
			It("errors", func() {
				fs.RegisterReadFileError("/path/ca1/certificate", errors.New("fake-error1"))

				_, err := subject.GetCertificatePEM()

				Expect(err).To(HaveOccurred())
			})
		})
	})
})
