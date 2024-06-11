<template>
    <v-container >
        <v-card elevation="4"  >
            <v-card-title>
                Login History
                <v-spacer></v-spacer>
            </v-card-title>
            <v-text-field v-model="search" append-outer-icon="mdi-magnify" label="Search" outlined hide-details
                rounded ></v-text-field>
            <v-data-table :color="transparent" :headers="headers" :items="loginHistory" :search="search" >
                <template slot="item.Date" slot-scope="{ item }">
                    {{ item.Date.toString().slice(0, 24) }}
                </template>
            </v-data-table>
        </v-card>
    </v-container>
</template>


<script>
import EventServices from '../services/EventServices';
export default {
    name: 'loginhistory',
    data() {
        return {
            search: '',
            loginHistory: []
        }
    },
    methods: {
        fetchLoginHistory() {
            EventServices.getHistory()
                .then(response => {
                    // console.log(response.data)
                    this.loginHistory = response.data.historyArr;
                })
                .catch(error => {
                    console.error('API Error:', error);
                });
        },
        arr() {
            return this.loginHistory;
        }
    },
    computed: {

        headers() {
            return [
                { text: 'Login Id', value: 'loginId' },
                { text: 'User Id', value: 'userId' },
                { text: 'Login Date', value: 'loginDate' },
                { text: 'Login Time', value: 'loginTime' },
                { text: 'Logout Date', value: 'loginDate' },
                { text: 'Logout Time', value: 'loginTime' },
            ]
        },

    },
    mounted() {
        this.fetchLoginHistory(); // Fetch login history when the component is mounted
    },
}
</script>
