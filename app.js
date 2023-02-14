const app = Vue.createApp({
    data() {
        return {
            totalResults: null,
            qstring: "",
            showMessage: false,
            singleMessage: null,
        }
    },
    methods: {
        search(qstring) {
            fetch('http://localhost:3000/api/search/'+qstring)
            .then(res => res.json())
            .then(data => this.totalResults = data)
            .catch(err => console.log(err.message))
        },
        allMessages(qstring) {
            this.totalResults  = this.search(qstring)
        },
        toggleMessage() {
            this.showMessage = !this.showMessage
        },
        message(id) {
            fetch('http://localhost:3000/api/search/_id:'+id)
            .then(res => res.json())
            .then(data => this.singleMessage = data)
            .catch(err => console.log(err.message))
            this.toggleMessage()
        }
    }
})

app.mount('#app')