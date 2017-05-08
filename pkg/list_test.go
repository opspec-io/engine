package pkg

import (
	"errors"
	"github.com/golang-interfaces/iioutil"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("List", func() {

	Context("when ioutil.ReadDir returns an error", func() {

		It("should be returned", func() {

			/* arrange */
			expectedError := errors.New("dummyError")

			fakeIOUtil := new(iioutil.Fake)
			fakeIOUtil.ReadDirReturns(nil, expectedError)

			objectUnderTest := pkg{
				ioUtil:               fakeIOUtil,
				manifestUnmarshaller: new(fakeManifestUnmarshaller),
			}

			/* act */
			_, actualError := objectUnderTest.List("")

			/* assert */
			Expect(actualError).To(Equal(expectedError))

		})

	})

})
