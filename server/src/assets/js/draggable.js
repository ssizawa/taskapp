import { VueElement } from 'vue';
import draggable from 'vuedraggable'

Vue.createApp({
    components:{
        'draggable': draggable,
    },
    data:{
        tasks1:[
            {item: 'task1'},
            {item: 'task3'}
        ],
        tasks2:[
            {item: 'task2'},
            {item: 'task4'}
        ]
    },
    computed: {
        task: function(){
            return this.tasks1;
        }
    }
}).mount('#app');