var vm = new Vue({
    data: {
        a: 1,
    },
    created: function () {
    // `this` указывает на экземпляр vm
    console.log('Значение a: ' + this.a)
    }
})