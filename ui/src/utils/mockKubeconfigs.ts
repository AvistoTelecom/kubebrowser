import type { Kubeconfig } from '@/types/Kubeconfig'

// List of real cluster names
const clusterNames = [
  'Internal Preprod',
  'Internal Prod',
  'External Dev',
  'External Prod',
  'External Preprod',
  'Staging',
  'QA',
  'Sandbox',
  'Customer Demo',
  'Legacy Cluster',
]

// Helper to generate a random base64 string of given length
function randomBase64(length: number): string {
  const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789='
  let result = ''
  for (let i = 0; i < length; i++) {
    result += chars.charAt(Math.floor(Math.random() * chars.length))
  }
  return result
}

export function generateMockKubeconfigs(count: number): Kubeconfig[] {
  return Array.from({ length: count }, (_, i) => {
    const caLength = 150 + Math.floor(Math.random() * 50)
    const idTokenLength = 100 + Math.floor(Math.random() * 200)
    const refreshTokenLength = 100 + Math.floor(Math.random() * 10)
    const clusterName = clusterNames[i % clusterNames.length]
    return {
      name: clusterName,
      kubeconfig: {
        apiVersion: 'v1',
        kind: 'Config',
        clusters: [
          {
            name: clusterName.replace(/\s+/g, '-').toLowerCase(),
            cluster: {
              server: `https://${clusterName.replace(/\s+/g, '-').toLowerCase()}.example.com`,
              'certificate-authority-data': randomBase64(caLength),
            },
          },
        ],
        contexts: [
          {
            name: `${clusterName.replace(/\s+/g, '-').toLowerCase()}-context`,
            context: {
              cluster: clusterName.replace(/\s+/g, '-').toLowerCase(),
              user: `user${i + 1}`,
            },
          },
        ],
        'current-context': `${clusterName.replace(/\s+/g, '-').toLowerCase()}-context`,
        users: [
          {
            name: `user${i + 1}`,
            user: {
              'auth-provider': {
                name: 'oidc',
                config: {
                  'client-id': `clientid${i + 1}`,
                  'client-secret': `secret${i + 1}`,
                  'id-token': randomBase64(idTokenLength),
                  'idp-issuer-url': `https://issuer${i + 1}.example.com`,
                  'refresh-token': randomBase64(refreshTokenLength),
                },
              },
            },
          },
        ],
      },
    }
  })
}
