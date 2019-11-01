window.onload = function () {


    var vm = new Vue({
        el: '#list-1',
        data: {
            items: [],
        },
        created: function () {
        // `this` указывает на экземпляр vm
        axios
            .get('http://localhost:8081/api/v1/authors')
            .then(response => (this.items = response["data"]));        
        }
    })


}