var vm = new Vue({
    data: {
        a: 1,
    },
    created: function () {
    // `this` указывает на экземпляр vm
    axios
        .get('http://localhost:8081/api/v1/authors')
        .then(response => (console.log(response["data"])));        
    }
})
