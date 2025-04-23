import axios from 'axios'

import type { Kubeconfig } from '@/types/Kubeconfig'
import { generateMockKubeconfigs } from '@/utils/mockKubeconfigs'

export async function getMe(): Promise<string> {
  if (import.meta.env.DEV) {
    // await new Promise((resolve) => setTimeout(resolve, 3000))
    // Mock response for development
    return 'Firstname Lastname'
  } else {
    // TODO: handle errors?
    return axios
      .get<string>('/api/me')
      .then((res) => res.data)
      .catch(() => '')
  }
}

export async function getConfigs(): Promise<Kubeconfig[]> {
  if (import.meta.env.DEV) {
    // await new Promise((resolve) => setTimeout(resolve, 1000))
    // Mock response for development
    return generateMockKubeconfigs(8)
  } else {
    // TODO: handle errors?
    return axios
      .get<Kubeconfig[]>('/api/kubeconfigs')
      .then((res) => res.data.sort((a, b) => a.name.localeCompare(b.name)))
      .catch(() => [])
  }
}
