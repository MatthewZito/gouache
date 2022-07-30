import { defineStore } from 'pinia'
import type { Resource } from '@/types'
import type { ResourceStoreState } from './types'
import { UserActionException } from '@/services/error'

/**
 * The resource store manages app-wide resources state.
 */
export const useResourceStore = defineStore('resource', {
  state: (): ResourceStoreState => ({
    /**
     * Currently loaded resources.
     */
    resources: [],
  }),

  actions: {
    setResources(resources: Resource[]) {
      this.resources = resources
    },

    patchResources(resource: Resource) {
      const targetIdx = this.resources.findIndex(({ id }) => id === resource.id)
      if (targetIdx === -1) {
        throw new UserActionException({
          internal: 'patchResources id === -1',
          friendly:
            'Something went wrong while updating your resources. Try refreshing the table or contact support.',
        })
      }

      this.resources[targetIdx] = resource
    },

    prependResource(resource: Resource) {
      this.resources = [resource, ...this.resources]
    },
  },
})
