window.onload = function () {


    var vm = new Vue({
        el: '#list-1',
        data: {
            items: [],
        },
        created: function () {        
            axios
                .get('http://localhost:8081/api/v1/authors')
                .then(response => (this.items = response["data"]));
        },
        methods : {
            removeList: function (item) {
                axios
                    .get('http://localhost:8081/api/v1/authors/'+item+'/albums')
                    .then(response => (this.items = response["data"]));
            }
        }
    })


}