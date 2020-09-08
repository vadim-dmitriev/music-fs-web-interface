window.onload = function () {


    var vm = new Vue({
        el: '#main',

        data: {
            mode: '',
            items: [],

            artist: '',
            album: '',

            isBtnBackShow: '',
        },

        created: function () {
            this.mode = 'artists'
            this.isBtnBackShow = 'none'
        },

        watch: {
            mode: function () {
                switch (this.mode) {
                    case 'artists':
                        this.showArtists();
                        // this.mode = 'albums';
                        break;
                    
                    case 'albums':
                        // this.artist = item;
                        this.showAlbums(this.artist);
                        this.isBtnBackShow = 'block'
                        // this.mode = 'songs';

                        break;
                    
                    case 'songs':
                        // this.album = item;
                        this.showSongs(this.album);
                        this.isBtnBackShow = 'block'
                        break;

                    default:
                        alert("Mode is unknown: '" + this.mode + "'")
                }
            },
        },

        methods : {
            next: function (item) {
                if (this.mode == 'artists') {
                    this.artist = item;
                    this.mode = 'albums';
                    return
                }
                if (this.mode == 'albums') {
                    this.album = item;
                    this.mode = 'songs';
                    return
                }
            },

            back: function () {
                if (this.mode == 'songs') {
                    this.album = '';
                    this.mode = 'albums';
                    return
                }
                if (this.mode == 'albums') {
                    this.artist = '';
                    this.mode = 'artists';
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

}