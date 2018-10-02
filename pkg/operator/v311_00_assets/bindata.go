// Code generated by go-bindata.
// sources:
// bindata/v3.11.0/kube-scheduler/cm.yaml
// bindata/v3.11.0/kube-scheduler/defaultconfig.yaml
// bindata/v3.11.0/kube-scheduler/deployment.yaml
// bindata/v3.11.0/kube-scheduler/ns.yaml
// bindata/v3.11.0/kube-scheduler/public-info-role.yaml
// bindata/v3.11.0/kube-scheduler/public-info-rolebinding.yaml
// bindata/v3.11.0/kube-scheduler/public-info.yaml
// bindata/v3.11.0/kube-scheduler/sa.yaml
// bindata/v3.11.0/kube-scheduler/scheduler-clusterrolebinding.yaml
// bindata/v3.11.0/kube-scheduler/svc.yaml
// DO NOT EDIT!

package v311_00_assets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _v3110KubeSchedulerCmYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  namespace: openshift-kube-scheduler
  name: deployment-scheduler-config
data:
  config.yaml:
`)

func v3110KubeSchedulerCmYamlBytes() ([]byte, error) {
	return _v3110KubeSchedulerCmYaml, nil
}

func v3110KubeSchedulerCmYaml() (*asset, error) {
	bytes, err := v3110KubeSchedulerCmYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-scheduler/cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeSchedulerDefaultconfigYaml = []byte(`apiVersion: componentconfig/v1alpha1
kind: KubeSchedulerConfiguration
`)

func v3110KubeSchedulerDefaultconfigYamlBytes() ([]byte, error) {
	return _v3110KubeSchedulerDefaultconfigYaml, nil
}

func v3110KubeSchedulerDefaultconfigYaml() (*asset, error) {
	bytes, err := v3110KubeSchedulerDefaultconfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-scheduler/defaultconfig.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeSchedulerDeploymentYaml = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: openshift-kube-scheduler
  name: scheduler
  labels:
    app: openshift-kube-scheduler
    scheduler: "true"
spec:
  strategy:
    type: RollingUpdate
  selector:
    matchLabels:
      app: openshift-kube-scheduler
      scheduler: "true"
  template:
    metadata:
      name: openshift-kube-scheduler
      labels:
        app: openshift-kube-scheduler
        scheduler: "true"
    spec:
      affinity:
        podAntiAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
          - labelSelector:
              matchLabels:
                app: openshift-kube-scheduler
            namespaces:
            - openshift-kube-scheduler
            topologyKey: kubernetes.io/hostname
      serviceAccountName: openshift-kube-scheduler-sa
      containers:
      - name: scheduler
        image: ${IMAGE}
        imagePullPolicy: IfNotPresent
        command: ["hyperkube", "kube-scheduler"]
        args:
        - "--config=/var/run/configmaps/config/config.yaml"
        ports:
        - containerPort: 8443
        volumeMounts:
        - mountPath: /var/run/configmaps/config
          name: config
        livenessProbe:
          httpGet:
            path: /healthz
            port: 10251
          initialDelaySeconds: 15
          timeoutSeconds: 15
      nodeSelector:
        node-role.kubernetes.io/master: ""
      tolerations:
      - operator: Exists
      volumes:
      - name: config
        configMap:
          name: deployment-scheduler-config
`)

func v3110KubeSchedulerDeploymentYamlBytes() ([]byte, error) {
	return _v3110KubeSchedulerDeploymentYaml, nil
}

func v3110KubeSchedulerDeploymentYaml() (*asset, error) {
	bytes, err := v3110KubeSchedulerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-scheduler/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeSchedulerNsYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  name: openshift-kube-scheduler
  labels:
    openshift.io/run-level: "0"
`)

func v3110KubeSchedulerNsYamlBytes() ([]byte, error) {
	return _v3110KubeSchedulerNsYaml, nil
}

func v3110KubeSchedulerNsYaml() (*asset, error) {
	bytes, err := v3110KubeSchedulerNsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-scheduler/ns.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeSchedulerPublicInfoRoleYaml = []byte(`kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: system:openshift:operator:kube-scheduler:public
rules:
- apiGroups:
  - storage.k8s.io
  resources:
  - storageclasses
  verbs:
  - get
  - list
  - watch
`)

func v3110KubeSchedulerPublicInfoRoleYamlBytes() ([]byte, error) {
	return _v3110KubeSchedulerPublicInfoRoleYaml, nil
}

func v3110KubeSchedulerPublicInfoRoleYaml() (*asset, error) {
	bytes, err := v3110KubeSchedulerPublicInfoRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-scheduler/public-info-role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeSchedulerPublicInfoRolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: system:openshift:operator:kube-scheduler:public
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:openshift:operator:kube-scheduler:public
subjects:
- kind: ServiceAccount
  name: openshift-kube-scheduler-sa
  namespace: openshift-kube-scheduler
`)

func v3110KubeSchedulerPublicInfoRolebindingYamlBytes() ([]byte, error) {
	return _v3110KubeSchedulerPublicInfoRolebindingYaml, nil
}

func v3110KubeSchedulerPublicInfoRolebindingYaml() (*asset, error) {
	bytes, err := v3110KubeSchedulerPublicInfoRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-scheduler/public-info-rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeSchedulerPublicInfoYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  namespace: openshift-kube-scheduler
  name: public-info
data:
  # version is the current of the kube-scheduler.  It is updated *after* it is being served consistently.
  version:
  # imagePolicyConfig.internalRegistryHostname is internal registry used for imagePolicyAdmission
  # TODO this probably won't make it to 4.0, we're likely to stuff the entire imagePolicyAdmission config in here
  imagePolicyConfig.internalRegistryHostname:
  # imagePolicyConfig.externalRegistryHostname is external registry used for imagePolicyAdmission
  # TODO this probably won't make it to 4.0, we're likely to stuff the entire imagePolicyAdmission config in here
  imagePolicyConfig.externalRegistryHostname:
  # defaultNodeSelector is used when no specific node selector is on a namespace
  # TODO we'd really like to see this collapsed onto upstream values
  projectConfig.defaultNodeSelector:
`)

func v3110KubeSchedulerPublicInfoYamlBytes() ([]byte, error) {
	return _v3110KubeSchedulerPublicInfoYaml, nil
}

func v3110KubeSchedulerPublicInfoYaml() (*asset, error) {
	bytes, err := v3110KubeSchedulerPublicInfoYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-scheduler/public-info.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeSchedulerSaYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: openshift-kube-scheduler
  name: openshift-kube-scheduler-sa
`)

func v3110KubeSchedulerSaYamlBytes() ([]byte, error) {
	return _v3110KubeSchedulerSaYaml, nil
}

func v3110KubeSchedulerSaYaml() (*asset, error) {
	bytes, err := v3110KubeSchedulerSaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-scheduler/sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeSchedulerSchedulerClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: system:openshift:operator:kube-scheduler:public-2
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:kube-scheduler
subjects:
- kind: ServiceAccount
  name: openshift-kube-scheduler-sa
  namespace: openshift-kube-scheduler
`)

func v3110KubeSchedulerSchedulerClusterrolebindingYamlBytes() ([]byte, error) {
	return _v3110KubeSchedulerSchedulerClusterrolebindingYaml, nil
}

func v3110KubeSchedulerSchedulerClusterrolebindingYaml() (*asset, error) {
	bytes, err := v3110KubeSchedulerSchedulerClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-scheduler/scheduler-clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeSchedulerSvcYaml = []byte(`apiVersion: v1
kind: Service
metadata:
  namespace: openshift-kube-scheduler
  name: scheduler
  annotations:
    service.alpha.openshift.io/serving-cert-secret-name: serving-cert
    prometheus.io/scrape: "true"
    prometheus.io/scheme: https
spec:
  selector:
    scheduler: "true"
  ports:
  - name: https
    port: 443
    targetPort: 8443
`)

func v3110KubeSchedulerSvcYamlBytes() ([]byte, error) {
	return _v3110KubeSchedulerSvcYaml, nil
}

func v3110KubeSchedulerSvcYaml() (*asset, error) {
	bytes, err := v3110KubeSchedulerSvcYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-scheduler/svc.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"v3.11.0/kube-scheduler/cm.yaml":                           v3110KubeSchedulerCmYaml,
	"v3.11.0/kube-scheduler/defaultconfig.yaml":                v3110KubeSchedulerDefaultconfigYaml,
	"v3.11.0/kube-scheduler/deployment.yaml":                   v3110KubeSchedulerDeploymentYaml,
	"v3.11.0/kube-scheduler/ns.yaml":                           v3110KubeSchedulerNsYaml,
	"v3.11.0/kube-scheduler/public-info-role.yaml":             v3110KubeSchedulerPublicInfoRoleYaml,
	"v3.11.0/kube-scheduler/public-info-rolebinding.yaml":      v3110KubeSchedulerPublicInfoRolebindingYaml,
	"v3.11.0/kube-scheduler/public-info.yaml":                  v3110KubeSchedulerPublicInfoYaml,
	"v3.11.0/kube-scheduler/sa.yaml":                           v3110KubeSchedulerSaYaml,
	"v3.11.0/kube-scheduler/scheduler-clusterrolebinding.yaml": v3110KubeSchedulerSchedulerClusterrolebindingYaml,
	"v3.11.0/kube-scheduler/svc.yaml":                          v3110KubeSchedulerSvcYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"v3.11.0": {nil, map[string]*bintree{
		"kube-scheduler": {nil, map[string]*bintree{
			"cm.yaml":                           {v3110KubeSchedulerCmYaml, map[string]*bintree{}},
			"defaultconfig.yaml":                {v3110KubeSchedulerDefaultconfigYaml, map[string]*bintree{}},
			"deployment.yaml":                   {v3110KubeSchedulerDeploymentYaml, map[string]*bintree{}},
			"ns.yaml":                           {v3110KubeSchedulerNsYaml, map[string]*bintree{}},
			"public-info-role.yaml":             {v3110KubeSchedulerPublicInfoRoleYaml, map[string]*bintree{}},
			"public-info-rolebinding.yaml":      {v3110KubeSchedulerPublicInfoRolebindingYaml, map[string]*bintree{}},
			"public-info.yaml":                  {v3110KubeSchedulerPublicInfoYaml, map[string]*bintree{}},
			"sa.yaml":                           {v3110KubeSchedulerSaYaml, map[string]*bintree{}},
			"scheduler-clusterrolebinding.yaml": {v3110KubeSchedulerSchedulerClusterrolebindingYaml, map[string]*bintree{}},
			"svc.yaml":                          {v3110KubeSchedulerSvcYaml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
