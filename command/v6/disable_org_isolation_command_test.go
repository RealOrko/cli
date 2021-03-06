package v6_test

import (
	"errors"

	"code.cloudfoundry.org/cli/actor/actionerror"
	"code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/command/commandfakes"
	. "code.cloudfoundry.org/cli/command/v6"
	"code.cloudfoundry.org/cli/command/v6/v6fakes"
	"code.cloudfoundry.org/cli/util/configv3"
	"code.cloudfoundry.org/cli/util/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var _ = Describe("disable-org-isolation Command", func() {
	var (
		cmd                  DisableOrgIsolationCommand
		testUI               *ui.UI
		fakeConfig           *commandfakes.FakeConfig
		fakeSharedActor      *commandfakes.FakeSharedActor
		fakeActor            *v6fakes.FakeDisableOrgIsolationActor
		binaryName           string
		executeErr           error
		isolationSegment     string
		org                  string
		deleteIsoSegWarnings v3action.Warnings
	)

	BeforeEach(func() {
		testUI = ui.NewTestUI(nil, NewBuffer(), NewBuffer())
		fakeConfig = new(commandfakes.FakeConfig)
		fakeSharedActor = new(commandfakes.FakeSharedActor)
		fakeActor = new(v6fakes.FakeDisableOrgIsolationActor)

		cmd = DisableOrgIsolationCommand{
			UI:          testUI,
			Config:      fakeConfig,
			SharedActor: fakeSharedActor,
			Actor:       fakeActor,
		}

		binaryName = "faceman"
		fakeConfig.BinaryNameReturns(binaryName)
		org = "org1"
		isolationSegment = "segment1"

		fakeConfig.CurrentUserReturns(configv3.User{Name: "admin"}, nil)
		cmd.RequiredArgs.OrganizationName = org
		cmd.RequiredArgs.IsolationSegmentName = isolationSegment

		deleteIsoSegWarnings = v3action.Warnings{"delete-isolation-segment-warning"}

		fakeActor.DeleteIsolationSegmentOrganizationByNameReturns(deleteIsoSegWarnings, nil)
	})

	JustBeforeEach(func() {
		executeErr = cmd.Execute(nil)
	})

	When("checking target fails", func() {
		BeforeEach(func() {
			fakeSharedActor.CheckTargetReturns(actionerror.NotLoggedInError{BinaryName: binaryName})
		})

		It("returns an error", func() {
			Expect(executeErr).To(MatchError(actionerror.NotLoggedInError{BinaryName: binaryName}))

			Expect(fakeSharedActor.CheckTargetCallCount()).To(Equal(1))
			checkTargetedOrg, checkTargetedSpace := fakeSharedActor.CheckTargetArgsForCall(0)
			Expect(checkTargetedOrg).To(BeFalse())
			Expect(checkTargetedSpace).To(BeFalse())
		})
	})

	When("user is not logged in", func() {
		BeforeEach(func() {
			fakeConfig.CurrentUserReturns(configv3.User{}, errors.New("user-not-logged-in"))
			cmd.RequiredArgs.OrganizationName = org
			cmd.RequiredArgs.IsolationSegmentName = isolationSegment
		})

		It("returns the error", func() {
			Expect(executeErr).To(MatchError("user-not-logged-in"))
		})
	})

	It("prints out that it will remove the entitlement", func() {
		Expect(testUI.Out).To(Say("Removing entitlement to isolation segment segment1 from org org1 as admin..."))
	})

	It("prints out warnings", func() {
		Expect(testUI.Err).To(Say(deleteIsoSegWarnings[0]))
	})

	It("Isolation segment is revoked from org", func() {
		Expect(executeErr).ToNot(HaveOccurred())
		Expect(testUI.Out).To(Say("OK"))

		Expect(fakeActor.DeleteIsolationSegmentOrganizationByNameCallCount()).To(Equal(1))
		actualIsolationSegmentName, actualOrgName := fakeActor.DeleteIsolationSegmentOrganizationByNameArgsForCall(0)
		Expect(actualIsolationSegmentName).To(Equal(isolationSegment))
		Expect(actualOrgName).To(Equal(org))
	})

	When("revoking fails", func() {
		BeforeEach(func() {
			fakeActor.DeleteIsolationSegmentOrganizationByNameReturns(deleteIsoSegWarnings, errors.New("delete failed boring message"))
		})

		It("returns an error and the warnings", func() {
			Expect(testUI.Err).To(Say(deleteIsoSegWarnings[0]))
			Expect(executeErr).To(MatchError("delete failed boring message"))
		})
	})
})
