<template>
<div id="app">
    <b-container>
        <img src="./assets/todoapp.png">
        <div>
            <TodoApp />
        </div>
    </b-container>
</div>
</template>

<script>
import TodoApp from './components/TodoApp.vue'

export default {
    components: {
        TodoApp
    }
}
</script>

<style>
#app {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 20px;
}
</style>
