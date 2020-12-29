<template>
  <div class="item">
    <b-list-group-item class="d-flex justify-content-between align-items-center">
      <span v-bind:class="{ 'task-done': task.done === true }">{{ task.name }}</span>
      <div>
        <b-button
          v-if="task.done === false"
          v-on:click="$parent.markTaskDone(task)"
          variant="outline-success"
          size="sm"
          class="border-0"
        >
          <b-icon-check2 font-scale="1.5"></b-icon-check2>
        </b-button>
        <b-button
          v-on:click="$parent.deleteTask(task)"
          variant="outline-danger"
          size="sm"
          class="border-0"
        >
          <b-icon-trash></b-icon-trash>
        </b-button>
      </div>
    </b-list-group-item>
  </div>
</template>

<script>
export default {
  props: ["task"],
};
</script>

<style>
.task-done {
  opacity: 0.8;
  text-decoration: line-through;
}
</style>
