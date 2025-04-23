export interface Kubeconfig {
  name: string
  kubeconfig: {
    apiVersion: string
    kind: string
    clusters: {
      name: string
      cluster: {
        server: string
        'certificate-authority-data'?: string
        'insecure-skip-tls-verify'?: boolean
      }
    }[]
    contexts?: {
      name: string
      context: {
        cluster: string
        user: string
      }
    }[]
    'current-context'?: string
    users?: {
      name: string
      user: {
        'auth-provider': {
          name: string
          config: {
            'client-id': string
            'client-secret': string
            'id-token': string
            'idp-issuer-url': string
            'refresh-token': string
          }
        }
      }
    }[]
  }
}
