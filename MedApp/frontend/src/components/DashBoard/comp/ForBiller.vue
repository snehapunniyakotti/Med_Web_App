<template>
    <v-container class="mt-12">
        <v-card elevation="3">
            <v-card-title>Your Total Sales</v-card-title>
            <v-card-text>{{ sales.today_total }}
                <v-icon v-if="sales.sale" class="success--text mb-2">mdi-arrow-up</v-icon>
                <v-icon v-else class="error--text">mdi-arrow-down</v-icon> {{ sales.percent }} %
            </v-card-text>
        </v-card>
    </v-container>
</template>

<script>
import EventServices from '@/services/EventServices'; // Update the path based on your project structure

export default {
    name: 'ForBiller',
    data() {
        return {
            apiData: null,  // to store the API response
            apiError: null, // to store any API error
            userId: this.$route.params.Id,
        };
    },
    props: {
        type: String,
        id: String,
    },
    computed: {

        sales() {
            if (this.apiData) {
                // console.log("API Data:", this.apiData); // Check if the data is available

                let today_total = this.apiData.arr.current_day_sales;
                let yesterday_total = this.apiData.arr.previous_day_sales;

                // console.log("Today Total:", today_total); // Check today_total value
                // console.log("Yesterday Total:", yesterday_total); // Check yesterday_total value

                let percentageChange;
                if (yesterday_total !== 0) {
                    percentageChange = Math.abs((today_total - yesterday_total) / yesterday_total) * 100;
                } else {
                    percentageChange = 0;
                }

                // console.log("Percentage Change:", percentageChange); // Check percentageChange value

                return {
                    today_total: today_total,
                    percent: percentageChange.toFixed(2),
                    sale: today_total >= yesterday_total
                };
            } else {
                return {
                    today_total: 0,
                    percent: 0,
                    sale: false
                };
            }
        },

    },
    mounted() {
        // console.log("user Id in Billber dashboard  "+this.userId)
        this.fetchSalesData();
    },
    methods: {
        fetchSalesData() {
            let obj = {
                userId: this.userId
            }
            let jsondata = JSON.stringify(obj)
            EventServices.getBillerDashboard(jsondata)
                .then(response => {
                    // console.log(response.data);
                    this.apiData = response.data;
                })
                .catch(error => {
                    console.error(" API error:", error);
                    this.apiError = error;
                });
        }
    }
};
</script>


























