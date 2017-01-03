package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/client/collaboration"
	"koding/remoteapi/client/compute_provider"
	"koding/remoteapi/client/data_dog"
	"koding/remoteapi/client/git_provider"
	"koding/remoteapi/client/github"
	"koding/remoteapi/client/j_account"
	"koding/remoteapi/client/j_api_token"
	"koding/remoteapi/client/j_combined_app_storage"
	"koding/remoteapi/client/j_compute_stack"
	"koding/remoteapi/client/j_credential"
	"koding/remoteapi/client/j_custom_partials"
	"koding/remoteapi/client/j_domain_alias"
	"koding/remoteapi/client/j_group"
	"koding/remoteapi/client/j_invitation"
	"koding/remoteapi/client/j_machine"
	"koding/remoteapi/client/j_name"
	"koding/remoteapi/client/j_password_recovery"
	"koding/remoteapi/client/j_proposed_domain"
	"koding/remoteapi/client/j_provisioner"
	"koding/remoteapi/client/j_proxy_filter"
	"koding/remoteapi/client/j_proxy_restriction"
	"koding/remoteapi/client/j_reward"
	"koding/remoteapi/client/j_reward_campaign"
	"koding/remoteapi/client/j_session"
	"koding/remoteapi/client/j_snapshot"
	"koding/remoteapi/client/j_stack_template"
	"koding/remoteapi/client/j_tag"
	"koding/remoteapi/client/j_team_invitation"
	"koding/remoteapi/client/j_url_alias"
	"koding/remoteapi/client/j_user"
	"koding/remoteapi/client/j_workspace"
	"koding/remoteapi/client/o_auth"
	"koding/remoteapi/client/s3"
	"koding/remoteapi/client/shared_machine"
	"koding/remoteapi/client/sidebar"
	"koding/remoteapi/client/social_channel"
	"koding/remoteapi/client/social_message"
	"koding/remoteapi/client/social_notification"
	"koding/remoteapi/client/system"
	"koding/remoteapi/client/tracker"
)

// Default koding HTTP client.
var Default = NewHTTPClient(nil)

// NewHTTPClient creates a new koding HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Koding {
	if formats == nil {
		formats = strfmt.Default
	}
	transport := httptransport.New("localhost", "/", []string{"http", "https"})
	return New(transport, formats)
}

// New creates a new koding client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Koding {
	cli := new(Koding)
	cli.Transport = transport

	cli.Collaboration = collaboration.New(transport, formats)

	cli.ComputeProvider = compute_provider.New(transport, formats)

	cli.DataDog = data_dog.New(transport, formats)

	cli.GitProvider = git_provider.New(transport, formats)

	cli.Github = github.New(transport, formats)

	cli.JAccount = j_account.New(transport, formats)

	cli.JAPIToken = j_api_token.New(transport, formats)

	cli.JCombinedAppStorage = j_combined_app_storage.New(transport, formats)

	cli.JComputeStack = j_compute_stack.New(transport, formats)

	cli.JCredential = j_credential.New(transport, formats)

	cli.JCustomPartials = j_custom_partials.New(transport, formats)

	cli.JDomainAlias = j_domain_alias.New(transport, formats)

	cli.JGroup = j_group.New(transport, formats)

	cli.JInvitation = j_invitation.New(transport, formats)

	cli.JMachine = j_machine.New(transport, formats)

	cli.JName = j_name.New(transport, formats)

	cli.JPasswordRecovery = j_password_recovery.New(transport, formats)

	cli.JProposedDomain = j_proposed_domain.New(transport, formats)

	cli.JProvisioner = j_provisioner.New(transport, formats)

	cli.JProxyFilter = j_proxy_filter.New(transport, formats)

	cli.JProxyRestriction = j_proxy_restriction.New(transport, formats)

	cli.JReward = j_reward.New(transport, formats)

	cli.JRewardCampaign = j_reward_campaign.New(transport, formats)

	cli.JSession = j_session.New(transport, formats)

	cli.JSnapshot = j_snapshot.New(transport, formats)

	cli.JStackTemplate = j_stack_template.New(transport, formats)

	cli.JTag = j_tag.New(transport, formats)

	cli.JTeamInvitation = j_team_invitation.New(transport, formats)

	cli.JURLAlias = j_url_alias.New(transport, formats)

	cli.JUser = j_user.New(transport, formats)

	cli.JWorkspace = j_workspace.New(transport, formats)

	cli.OAuth = o_auth.New(transport, formats)

	cli.S3 = s3.New(transport, formats)

	cli.SharedMachine = shared_machine.New(transport, formats)

	cli.Sidebar = sidebar.New(transport, formats)

	cli.SocialChannel = social_channel.New(transport, formats)

	cli.SocialMessage = social_message.New(transport, formats)

	cli.SocialNotification = social_notification.New(transport, formats)

	cli.System = system.New(transport, formats)

	cli.Tracker = tracker.New(transport, formats)

	return cli
}

// Koding is a client for koding
type Koding struct {
	Collaboration *collaboration.Client

	ComputeProvider *compute_provider.Client

	DataDog *data_dog.Client

	GitProvider *git_provider.Client

	Github *github.Client

	JAccount *j_account.Client

	JAPIToken *j_api_token.Client

	JCombinedAppStorage *j_combined_app_storage.Client

	JComputeStack *j_compute_stack.Client

	JCredential *j_credential.Client

	JCustomPartials *j_custom_partials.Client

	JDomainAlias *j_domain_alias.Client

	JGroup *j_group.Client

	JInvitation *j_invitation.Client

	JMachine *j_machine.Client

	JName *j_name.Client

	JPasswordRecovery *j_password_recovery.Client

	JProposedDomain *j_proposed_domain.Client

	JProvisioner *j_provisioner.Client

	JProxyFilter *j_proxy_filter.Client

	JProxyRestriction *j_proxy_restriction.Client

	JReward *j_reward.Client

	JRewardCampaign *j_reward_campaign.Client

	JSession *j_session.Client

	JSnapshot *j_snapshot.Client

	JStackTemplate *j_stack_template.Client

	JTag *j_tag.Client

	JTeamInvitation *j_team_invitation.Client

	JURLAlias *j_url_alias.Client

	JUser *j_user.Client

	JWorkspace *j_workspace.Client

	OAuth *o_auth.Client

	S3 *s3.Client

	SharedMachine *shared_machine.Client

	Sidebar *sidebar.Client

	SocialChannel *social_channel.Client

	SocialMessage *social_message.Client

	SocialNotification *social_notification.Client

	System *system.Client

	Tracker *tracker.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Koding) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Collaboration.SetTransport(transport)

	c.ComputeProvider.SetTransport(transport)

	c.DataDog.SetTransport(transport)

	c.GitProvider.SetTransport(transport)

	c.Github.SetTransport(transport)

	c.JAccount.SetTransport(transport)

	c.JAPIToken.SetTransport(transport)

	c.JCombinedAppStorage.SetTransport(transport)

	c.JComputeStack.SetTransport(transport)

	c.JCredential.SetTransport(transport)

	c.JCustomPartials.SetTransport(transport)

	c.JDomainAlias.SetTransport(transport)

	c.JGroup.SetTransport(transport)

	c.JInvitation.SetTransport(transport)

	c.JMachine.SetTransport(transport)

	c.JName.SetTransport(transport)

	c.JPasswordRecovery.SetTransport(transport)

	c.JProposedDomain.SetTransport(transport)

	c.JProvisioner.SetTransport(transport)

	c.JProxyFilter.SetTransport(transport)

	c.JProxyRestriction.SetTransport(transport)

	c.JReward.SetTransport(transport)

	c.JRewardCampaign.SetTransport(transport)

	c.JSession.SetTransport(transport)

	c.JSnapshot.SetTransport(transport)

	c.JStackTemplate.SetTransport(transport)

	c.JTag.SetTransport(transport)

	c.JTeamInvitation.SetTransport(transport)

	c.JURLAlias.SetTransport(transport)

	c.JUser.SetTransport(transport)

	c.JWorkspace.SetTransport(transport)

	c.OAuth.SetTransport(transport)

	c.S3.SetTransport(transport)

	c.SharedMachine.SetTransport(transport)

	c.Sidebar.SetTransport(transport)

	c.SocialChannel.SetTransport(transport)

	c.SocialMessage.SetTransport(transport)

	c.SocialNotification.SetTransport(transport)

	c.System.SetTransport(transport)

	c.Tracker.SetTransport(transport)

}
