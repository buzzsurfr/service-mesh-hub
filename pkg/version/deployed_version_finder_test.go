package version_test

import (
	"context"
	"reflect"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/testutils"
	mock_kubernetes_apps "github.com/solo-io/mesh-projects/pkg/clients/kubernetes/apps/mocks"
	"github.com/solo-io/mesh-projects/pkg/common/docker"
	mock_docker "github.com/solo-io/mesh-projects/pkg/common/docker/mocks"
	"github.com/solo-io/mesh-projects/pkg/env"
	"github.com/solo-io/mesh-projects/pkg/version"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("Federation Decider", func() {
	var (
		ctrl                       *gomock.Controller
		ctx                        context.Context
		containerAndDeploymentName = "mesh-networking"
		testErr                    = eris.New("test-err")
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.TODO()
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("open source version", func() {
		It("can find the open source version based on the presence of mesh-networking", func() {
			imageNameParser := mock_docker.NewMockImageNameParser(ctrl)
			deploymentClient := mock_kubernetes_apps.NewMockDeploymentClient(ctrl)

			deployedVersionFinder := version.NewDeployedVersionFinder(deploymentClient, imageNameParser)

			deploymentClient.EXPECT().
				Get(ctx, client.ObjectKey{
					Name:      containerAndDeploymentName,
					Namespace: env.DefaultWriteNamespace,
				}).
				Return(&v1.Deployment{
					Spec: v1.DeploymentSpec{
						Template: v12.PodTemplateSpec{
							Spec: v12.PodSpec{
								Containers: []v12.Container{{
									Name:  containerAndDeploymentName,
									Image: "mesh-networking-image-name",
								}},
							},
						},
					},
				}, nil)

			imageNameParser.EXPECT().
				Parse("mesh-networking-image-name").
				Return(&docker.Image{
					Tag: "1.0.0",
				}, nil)

			openSourceVersion, err := deployedVersionFinder.OpenSourceVersion(ctx, env.DefaultWriteNamespace)
			Expect(err).NotTo(HaveOccurred())
			Expect(openSourceVersion).To(Equal("1.0.0"))
		})

		It("can return an open source version with a leading 'v' stripped", func() {
			imageNameParser := mock_docker.NewMockImageNameParser(ctrl)
			deploymentClient := mock_kubernetes_apps.NewMockDeploymentClient(ctrl)

			deployedVersionFinder := version.NewDeployedVersionFinder(deploymentClient, imageNameParser)

			deploymentClient.EXPECT().
				Get(ctx, client.ObjectKey{
					Name:      containerAndDeploymentName,
					Namespace: env.DefaultWriteNamespace,
				}).
				Return(&v1.Deployment{
					Spec: v1.DeploymentSpec{
						Template: v12.PodTemplateSpec{
							Spec: v12.PodSpec{
								Containers: []v12.Container{{
									Name:  containerAndDeploymentName,
									Image: "mesh-networking-image-name",
								}},
							},
						},
					},
				}, nil)

			imageNameParser.EXPECT().
				Parse("mesh-networking-image-name").
				Return(&docker.Image{
					Tag: "v1.0.0",
				}, nil)

			openSourceVersion, err := deployedVersionFinder.OpenSourceVersion(ctx, env.DefaultWriteNamespace)
			Expect(err).NotTo(HaveOccurred())
			Expect(openSourceVersion).To(Equal("1.0.0"))
		})

		It("responds with the appropriate error if open-source SMH is not installed", func() {
			imageNameParser := mock_docker.NewMockImageNameParser(ctrl)
			deploymentClient := mock_kubernetes_apps.NewMockDeploymentClient(ctrl)

			deployedVersionFinder := version.NewDeployedVersionFinder(deploymentClient, imageNameParser)

			deploymentClient.EXPECT().
				Get(ctx, client.ObjectKey{
					Name:      containerAndDeploymentName,
					Namespace: env.DefaultWriteNamespace,
				}).
				Return(nil, errors.NewNotFound(schema.GroupResource{}, "test-resource-name"))

			openSourceVersion, err := deployedVersionFinder.OpenSourceVersion(ctx, env.DefaultWriteNamespace)
			Expect(err).To(testutils.HaveInErrorChain(version.NoOpenSourceDeployment))
			Expect(openSourceVersion).To(BeEmpty())
		})

		It("responds with the appropriate error if deployment lookup fails", func() {
			imageNameParser := mock_docker.NewMockImageNameParser(ctrl)
			deploymentClient := mock_kubernetes_apps.NewMockDeploymentClient(ctrl)

			deployedVersionFinder := version.NewDeployedVersionFinder(deploymentClient, imageNameParser)
			reflect.DeepEqual(map[string]string{}, map[string]string{})

			deploymentClient.EXPECT().
				Get(ctx, client.ObjectKey{
					Name:      containerAndDeploymentName,
					Namespace: env.DefaultWriteNamespace,
				}).
				Return(nil, testErr)

			openSourceVersion, err := deployedVersionFinder.OpenSourceVersion(ctx, env.DefaultWriteNamespace)
			Expect(err).To(testutils.HaveInErrorChain(version.FailedToLookUpOpenSourceDeployemnt(testErr)))
			Expect(openSourceVersion).To(BeEmpty())
		})

		It("responds with the appropriate error if the expected container is missing", func() {
			imageNameParser := mock_docker.NewMockImageNameParser(ctrl)
			deploymentClient := mock_kubernetes_apps.NewMockDeploymentClient(ctrl)

			deployedVersionFinder := version.NewDeployedVersionFinder(deploymentClient, imageNameParser)

			deploymentClient.EXPECT().
				Get(ctx, client.ObjectKey{
					Name:      containerAndDeploymentName,
					Namespace: env.DefaultWriteNamespace,
				}).
				Return(&v1.Deployment{
					Spec: v1.DeploymentSpec{
						Template: v12.PodTemplateSpec{
							Spec: v12.PodSpec{
								Containers: []v12.Container{{
									Name:  "intentionally-wrong-container-name",
									Image: "mesh-networking-image-name",
								}},
							},
						},
					},
				}, nil)

			openSourceVersion, err := deployedVersionFinder.OpenSourceVersion(ctx, env.DefaultWriteNamespace)
			Expect(err).To(testutils.HaveInErrorChain(version.FailedToFindContainer(containerAndDeploymentName, containerAndDeploymentName)))
			Expect(openSourceVersion).To(BeEmpty())
		})
	})
})
