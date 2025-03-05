// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/openshift/api/config/v1"
	configv1 "github.com/openshift/client-go/config/applyconfigurations/config/v1"
	typedconfigv1 "github.com/openshift/client-go/config/clientset/versioned/typed/config/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeOAuths implements OAuthInterface
type fakeOAuths struct {
	*gentype.FakeClientWithListAndApply[*v1.OAuth, *v1.OAuthList, *configv1.OAuthApplyConfiguration]
	Fake *FakeConfigV1
}

func newFakeOAuths(fake *FakeConfigV1) typedconfigv1.OAuthInterface {
	return &fakeOAuths{
		gentype.NewFakeClientWithListAndApply[*v1.OAuth, *v1.OAuthList, *configv1.OAuthApplyConfiguration](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("oauths"),
			v1.SchemeGroupVersion.WithKind("OAuth"),
			func() *v1.OAuth { return &v1.OAuth{} },
			func() *v1.OAuthList { return &v1.OAuthList{} },
			func(dst, src *v1.OAuthList) { dst.ListMeta = src.ListMeta },
			func(list *v1.OAuthList) []*v1.OAuth { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.OAuthList, items []*v1.OAuth) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
