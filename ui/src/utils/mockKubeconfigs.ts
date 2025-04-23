import type { Kubeconfig } from '@/types/Kubeconfig'

// Helper to generate a random base64 string of given length
function randomBase64(length: number): string {
  const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/='
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
    return {
      name: `Cluster number ${i + 1}`,
      kubeconfig: {
        apiVersion: 'v1',
        kind: 'Config',
        clusters: [
          {
            name: `cluster${i + 1}`,
            cluster: {
              server: `https://cluster${i + 1}.example.com`,
              'certificate-authority-data': randomBase64(caLength),
            },
          },
        ],
        contexts: [
          {
            name: `context${i + 1}`,
            context: {
              cluster: `cluster${i + 1}`,
              user: `user${i + 1}`,
            },
          },
        ],
        'current-context': `context${i + 1}`,
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
