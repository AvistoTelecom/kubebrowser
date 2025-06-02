load('ext://helm_resource', 'helm_resource')
load('ext://ko', 'ko_build')
watch_file('./server')

def get_kubeconfigs_names():
  return ['kubeconfig-' + str(i) for i in range(1, 6)]

def get_kubeconfigs(folder):
  return [folder + '/' + name + '.yaml' for name in get_kubeconfigs_names()]

ko_build(ref='server',
        import_path='./server',
        deps=['./server']
)

kubeconfigs = get_kubeconfigs('./dev/kubeconfigs/')

local_resource('frontend build', cmd='./scripts/buildFrontForBack.sh', deps="./ui/src")

helm_resource(
  'server',
  'avisto/kubebrowser',
  namespace='default',
  deps=['./dev/values.yaml'],
  image_deps=['server'],
  port_forwards=[8080],
  image_keys=[('server.image.registry', 'server.image.repository', 'server.image.tag')],
  flags=['--values=./dev/values.yaml'])

k8s_resource(new_name='kubeconfigs crds', resource_deps=['server'], objects=get_kubeconfigs_names())
k8s_yaml(kubeconfigs)
