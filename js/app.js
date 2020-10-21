'use strict';

// need to be in global context
const getToDo = () => {
    return fetch('./json/todo_list.json');
}
