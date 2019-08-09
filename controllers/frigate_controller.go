/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	shipv1beta1 "github.com/weizhang555/test/api/v1beta1"
)

// FrigateReconciler reconciles a Frigate object
type FrigateReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=ship.k8s.io,resources=frigates,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ship.k8s.io,resources=frigates/status,verbs=get;update;patch

func (r *FrigateReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	//log := r.Log.WithValues("frigate", req.NamespacedName)
	log := r.Log

	// your logic here
	listOptFunc := func(opt *client.ListOptions) {
		opt.Namespace = "default"
	}

	podList := &v1.PodList{}
	r.Client.List(ctx, podList, listOptFunc)

	for _, pod := range podList.Items {
		log.Info("pod", "name", pod.Name)
	}
	return ctrl.Result{}, nil
}

func (r *FrigateReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&shipv1beta1.Frigate{}).
		Complete(r)
}
