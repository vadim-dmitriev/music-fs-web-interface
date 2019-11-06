window.onload = function () {


    var vm = new Vue({
        el: '#list-1',

        data: {
            artist: '',
            album: '',
            mode: 'None',
            items: [],
        },

        created: function () {        
            this.show()
        },

        updated: function () {
            console.log(this.mode);
        },

        methods : {

            show: function (item) {
                switch (this.mode) {
                    case 'None':
                        this.showArtists();
                        this.mode = 'artists';
                        btnBack.isShow = 'none'
                        break;
                    
                    case 'artists':
                        this.artist = item;
                        this.showAlbums(item);
                        this.mode = 'albums';
                        btnBack.isShow = 'block'
                        break;
                    
                    case 'albums':
                        this.album = item;
                        this.showSongs(item);
                        this.mode = 'songs';
                        btnBack.isShow = 'block'
                        break;

                    default:
                        alert("Mode is unknown: '" + this.mode + "'")
                }
            },

            showArtists: function () {
                axios
                    .get('http://localhost:8081/api/v1/authors')
                    .then(response => (this.items = response["data"]));
            },

            showAlbums: function (artist) {
                axios
                    .get('http://localhost:8081/api/v1/authors/'+this.artist+'/albums')
                    .then(response => (this.items = response["data"]));
            },

            showSongs: function (album) {
                axios
                    .get('http://localhost:8081/api/v1/authors/'+this.artist+'/albums/' + this.album + '/songs')
                    .then(response => (this.items = response["data"]));
            }

        }
    })


    var btnBack = new Vue({
        el: '#btnBack',
        data: {
            // isShow: this.parent.mode == 'artists' || this.parent.mode == 'None' ? 'none' : 'block'
            isShow: ''
        },
        methods: {
            back: function () {
                vm.showArtists()
            }
        }
    })


}