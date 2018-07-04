// root shared state

import { addDefaultMutations } from "app/util"

const state = {
  todoList: [],
}

const getters = {}

const mutations = {
  addTodo(state, s) {
    state.todoList.push({ value: s.todo })
  },
}

const actions = {}

addDefaultMutations(actions, mutations)
export default {
  state,
  getters,
  actions,
  mutations,
  namespaced: true,
}
