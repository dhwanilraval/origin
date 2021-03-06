// This file was automatically generated by informer-gen

package v1

import (
	oauth_v1 "github.com/openshift/origin/pkg/oauth/apis/oauth/v1"
	clientset "github.com/openshift/origin/pkg/oauth/generated/clientset"
	internalinterfaces "github.com/openshift/origin/pkg/oauth/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/origin/pkg/oauth/generated/listers/oauth/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// OAuthClientAuthorizationInformer provides access to a shared informer and lister for
// OAuthClientAuthorizations.
type OAuthClientAuthorizationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.OAuthClientAuthorizationLister
}

type oAuthClientAuthorizationInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newOAuthClientAuthorizationInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.OauthV1().OAuthClientAuthorizations().List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.OauthV1().OAuthClientAuthorizations().Watch(options)
			},
		},
		&oauth_v1.OAuthClientAuthorization{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *oAuthClientAuthorizationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&oauth_v1.OAuthClientAuthorization{}, newOAuthClientAuthorizationInformer)
}

func (f *oAuthClientAuthorizationInformer) Lister() v1.OAuthClientAuthorizationLister {
	return v1.NewOAuthClientAuthorizationLister(f.Informer().GetIndexer())
}
