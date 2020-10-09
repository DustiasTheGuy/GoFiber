import axios from "axios";

export default {
    name: 'AddComponent',
    props: {},
    data() {
        return {
            name: undefined,
            age: undefined,
            department: undefined,
            salary: undefined,
            message: undefined,
            msgColor: "green",
            showSnackbar: false
        }
    },
    methods: {
        submit() {
            let data = {
                name: this.name,
                age: parseInt(this.age),
                department: this.department,
                salary: parseInt(this.salary)
            }
            axios.post("http://localhost:3000/create", data)
                .then(response => {
                    console.log(response)
                    this.showSnackbar = true
                    this.message = response.data.message

                    if (response.data.success) {
                        this.msgColor = "#006e10"
                        this.name = undefined
                        this.age = undefined
                        this.department = undefined
                        this.salary = undefined
                    } else {
                        this.msgColor = "#ff2929"
                    }
                })
                .catch(err => {
                    console.log(err)
                })

        }
    }
}