package core

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opspec-io/sdk-golang/model"
	"path/filepath"
)

var _ = Context("caller", func() {
	Context("newCaller", func() {
		It("should return caller", func() {
			/* arrange/act/assert */
			Expect(newCaller(new(fakeContainerCaller))).To(Not(BeNil()))
		})
	})
	Context("Call", func() {
		Context("Null SCG", func() {
			It("should not throw", func() {
				/* arrange */
				fakeContainerCaller := new(fakeContainerCaller)

				/* act */
				objectUnderTest := newCaller(fakeContainerCaller)

				/* assert */
				objectUnderTest.Call(
					"dummyCallId",
					map[string]*model.Data{},
					nil,
					"dummyPkgRef",
					"dummyRootOpId",
				)
			})
		})
		Context("Container SCG", func() {
			It("should call containerCaller.Call w/ expected args", func() {
				/* arrange */
				fakeContainerCaller := new(fakeContainerCaller)

				providedCallId := "dummyCallId"
				providedArgs := map[string]*model.Data{}
				providedSCG := &model.SCG{
					Container: &model.SCGContainerCall{},
				}
				providedPkgRef := "dummyPkgRef"
				providedRootOpId := "dummyRootOpId"

				objectUnderTest := newCaller(fakeContainerCaller)

				/* act */
				objectUnderTest.Call(
					providedCallId,
					providedArgs,
					providedSCG,
					providedPkgRef,
					providedRootOpId,
				)

				/* assert */
				actualArgs,
					actualCallId,
					actualSCG,
					actualPkgRef,
					actualRootOpId := fakeContainerCaller.CallArgsForCall(0)
				Expect(actualCallId).To(Equal(providedCallId))
				Expect(actualArgs).To(Equal(providedArgs))
				Expect(actualSCG).To(Equal(providedSCG.Container))
				Expect(actualPkgRef).To(Equal(providedPkgRef))
				Expect(actualRootOpId).To(Equal(providedRootOpId))
			})
		})
		Context("Op SCG", func() {
			It("should call opCaller.Call w/ expected args", func() {
				/* arrange */
				fakeOpCaller := new(fakeOpCaller)

				providedCallId := "dummyCallId"
				providedArgs := map[string]*model.Data{}
				providedSCG := &model.SCG{
					Op: &model.SCGOpCall{
						Pkg: &model.SCGOpCallPkg{
							Ref: "dummyPkgRef",
						},
					},
				}
				providedPkgRef := "dummyPkgRef"
				providedRootOpId := "dummyRootOpId"

				objectUnderTest := newCaller(new(fakeContainerCaller))
				objectUnderTest.setOpCaller(fakeOpCaller)

				/* act */
				objectUnderTest.Call(
					providedCallId,
					providedArgs,
					providedSCG,
					providedPkgRef,
					providedRootOpId,
				)

				/* assert */
				actualArgs,
					actualCallId,
					actualPkgRef,
					actualRootOpId,
					actualSCG := fakeOpCaller.CallArgsForCall(0)
				Expect(actualCallId).To(Equal(providedCallId))
				Expect(actualArgs).To(Equal(providedArgs))
				Expect(actualPkgRef).To(Equal(filepath.Dir(providedPkgRef)))
				Expect(actualRootOpId).To(Equal(providedRootOpId))
				Expect(actualSCG).To(Equal(providedSCG.Op))
			})
		})
		Context("Parallel SCG", func() {
			It("should call parallelCaller.Call w/ expected args", func() {
				/* arrange */
				fakeParallelCaller := new(fakeParallelCaller)

				providedCallId := "dummyCallId"
				providedArgs := map[string]*model.Data{}
				providedSCG := &model.SCG{
					Parallel: []*model.SCG{
						{Container: &model.SCGContainerCall{}},
					},
				}
				providedPkgRef := "dummyPkgRef"
				providedRootOpId := "dummyRootOpId"

				objectUnderTest := newCaller(new(fakeContainerCaller))
				objectUnderTest.setParallelCaller(fakeParallelCaller)

				/* act */
				objectUnderTest.Call(
					providedCallId,
					providedArgs,
					providedSCG,
					providedPkgRef,
					providedRootOpId,
				)

				/* assert */
				providedCallId,
					actualArgs,
					actualRootOpId,
					actualPkgRef,
					actualSCG := fakeParallelCaller.CallArgsForCall(0)
				Expect(actualArgs).To(Equal(providedArgs))
				Expect(actualRootOpId).To(Equal(providedRootOpId))
				Expect(actualPkgRef).To(Equal(providedPkgRef))
				Expect(actualSCG).To(Equal(providedSCG.Parallel))
			})
		})
		Context("Serial SCG", func() {

			It("should call serialCaller.Call w/ expected args", func() {
				/* arrange */
				fakeSerialCaller := new(fakeSerialCaller)

				providedCallId := "dummyCallId"
				providedArgs := map[string]*model.Data{}
				providedSCG := &model.SCG{
					Serial: []*model.SCG{
						{Container: &model.SCGContainerCall{}},
					},
				}
				providedPkgRef := "dummyPkgRef"
				providedRootOpId := "dummyRootOpId"

				objectUnderTest := newCaller(new(fakeContainerCaller))
				objectUnderTest.setSerialCaller(fakeSerialCaller)

				/* act */
				objectUnderTest.Call(
					providedCallId,
					providedArgs,
					providedSCG,
					providedPkgRef,
					providedRootOpId,
				)

				/* assert */
				actualCallId,
					actualArgs,
					actualRootOpId,
					actualPkgRef,
					actualSCG := fakeSerialCaller.CallArgsForCall(0)
				Expect(actualCallId).To(Equal(providedCallId))
				Expect(actualArgs).To(Equal(providedArgs))
				Expect(actualRootOpId).To(Equal(providedRootOpId))
				Expect(actualPkgRef).To(Equal(providedPkgRef))
				Expect(actualSCG).To(Equal(providedSCG.Serial))
			})
		})
		Context("No SCG", func() {
			It("should error", func() {
				/* arrange */
				fakeSerialCaller := new(fakeSerialCaller)

				providedCallId := "dummyCallId"
				providedArgs := map[string]*model.Data{}
				providedSCG := &model.SCG{}
				providedPkgRef := "dummyPkgRef"
				providedRootOpId := "dummyRootOpId"
				expectedError := fmt.Errorf("Invalid call graph %+v\n", providedSCG)

				objectUnderTest := newCaller(new(fakeContainerCaller))
				objectUnderTest.setSerialCaller(fakeSerialCaller)

				/* act */
				actualError := objectUnderTest.Call(
					providedCallId,
					providedArgs,
					providedSCG,
					providedPkgRef,
					providedRootOpId,
				)

				/* assert */
				Expect(actualError).To(Equal(expectedError))
			})
		})
	})
})
