import axios from 'axios'

import type { Kubeconfig } from '@/types/Kubeconfig'

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
    // await new Promise((resolve) => setTimeout(resolve, 4000))
    // Mock response for development
    return [
      { name: 'Cluster number 1', kubeconfig: { apiVersion: 'v1', kind: 'Config' } },
      { name: 'Cluster number 2', kubeconfig: { apiVersion: 'v1', kind: 'Config2' } },
      { name: 'Cluster number 3', kubeconfig: { apiVersion: 'v1', kind: 'Config2' } },
      { name: 'Cluster number 4', kubeconfig: { apiVersion: 'v1', kind: 'Config2' } },
      { name: 'Cluster number 5', kubeconfig: { apiVersion: 'v1', kind: 'Config2' } },
      { name: 'Cluster number 6', kubeconfig: { apiVersion: 'v1', kind: 'Config2' } },
      { name: 'Cluster number 7', kubeconfig: { apiVersion: 'v1', kind: 'Config2' } },
      { name: 'Another cluster', kubeconfig: { apiVersion: 'v1', kind: 'Another' } },
    ]
  } else {
    // TODO: handle errors?
    return axios
      .get<Kubeconfig[]>('/api/kubeconfigs')
      .then((res) => res.data)
      .catch(() => [])
  }
}
