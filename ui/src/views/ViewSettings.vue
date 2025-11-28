<script setup lang="ts">
import IconChevronRight from '@/icons/IconChevronRight.vue'
import { ref, computed } from 'vue'

const verbs = ref(['*', 'create', 'update', 'delete', 'read'])

const permissions = ref({
  permissions: [
    {
      user: 'bidule',
      rules: [
        {
          resource: 'pods',
          verbs: ['delete', 'create'],
        },
        {
          resource: 'deployments',
          verbs: ['update'],
        },
      ],
    },
    {
      group: 'admin',
      rules: [
        {
          resource: '*',
          verbs: ['*'],
        },
      ],
    },
  ],
})

const totalCols = computed(() => verbs.value.length + 1)

const expanded = ref<Record<number, boolean>>({})

function toggle(index: number) {
  expanded.value[index] = !expanded.value[index]
}

function ruleHasVerb(ruleVerbs: string[] | undefined, verb: string) {
  if (!ruleVerbs) return false
  if (ruleVerbs.includes('*')) return true
  return ruleVerbs.includes(verb)
}
</script>

<template>
  <div class="mx-4 md:mx-8 py-2 text-xl font-extrabold md:py-4 md:text-2xl 2xl:py-8 2xl:text-3xl">
    <div>Admin settings</div>
  </div>

  <table class="min-w-full border-collapse">
    <thead>
      <tr>
        <th class="border-y border-l border-white p-2 text-left">Entité / Règles</th>
        <th v-for="verb in verbs" :key="verb" class="p-2 border-x border-t text-center">
          {{ verb }}
        </th>
      </tr>
    </thead>

    <tbody>
      <template v-for="(permission, index) in permissions.permissions" :key="index">
        <tr class="cursor-pointer" @click="toggle(index)" :aria-expanded="!!expanded[index]">
          <td
            class="border-x border-white p-2 font-medium"
            :colspan="totalCols"
            :class="{ 'border-b': !expanded[index] }"
          >
            <div class="flex items-center justify-between">
              <div>
                <span class="font-bold">{{ permission.user ? 'User' : 'Group' }}</span>
                {{ permission.user ?? permission.group }}
              </div>
              <div class="text-sm text-white w-6 h-6">
                <span v-if="expanded[index]"
                  ><IconChevronRight class="shrink-0 w-full h-full transform rotate-90"
                /></span>
                <span v-else><IconChevronRight class="shrink-0 w-full h-full" /></span>
              </div>
            </div>
          </td>
        </tr>

        <tr
          v-for="(rule, rIndex) in permission.rules"
          :key="`rule-${index}-${rIndex}`"
          v-show="expanded[index]"
        >
          <td class="p-2 border text-left">
            <div class="text-sm font-medium pl-4">{{ rule.resource }}</div>
          </td>

          <td
            v-for="verb in verbs"
            :key="verb"
            class="p-2 text-center align-middle border-x border-b"
          >
            <span v-if="ruleHasVerb(rule.verbs, verb)">✓</span>
          </td>
        </tr>

        <tr v-if="!(permission.rules && permission.rules.length) && expanded[index]">
          <td :colspan="totalCols" class="p-2 border text-sm text-gray-500 italic">
            No rules defined
          </td>
        </tr>
      </template>
    </tbody>
  </table>
</template>

<style scoped>
/* small visual polish: collapse border gaps */
table {
  border-spacing: 0;
}
</style>
