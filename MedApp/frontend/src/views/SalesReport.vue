<template>
    <v-container class="mt-12">
        <v-card elevation="0">
            <v-card-title class="text-uppercase">Sales Report</v-card-title>
            <v-form ref="form">
                <v-row justify="center">
                    <v-col cols="3">
                        <v-menu ref="fromDateMenu" v-model="fromDateMenu" :close-on-content-click="false"
                            :nudge-right="40" transition="scale-transition" offset-y>
                            <template v-slot:activator="{ on }">
                                <v-text-field v-model="fromDate" label="From " prepend-inner-icon="mdi-calendar"
                                    outlined v-on="on"></v-text-field>
                            </template>
                            <v-date-picker v-model="fromDate"
                                @input="$refs.fromDateMenu.save(fromDate)"></v-date-picker>
                        </v-menu>
                    </v-col>
                    <v-col cols="3">
                        <v-menu ref="toDateMenu" v-model="toDateMenu" :close-on-content-click="false" :nudge-right="40"
                            transition="scale-transition" offset-y>
                            <template v-slot:activator="{ on }">
                                <v-text-field v-model="toDate" label="To" v-on="on" outlined
                                    prepend-inner-icon="mdi-calendar"></v-text-field>
                            </template>
                            <v-date-picker v-model="toDate" @input="$refs.toDateMenu.save(toDate)"></v-date-picker>
                        </v-menu>
                    </v-col>
                    <v-col cols="3">
                        <v-btn class="ml-2 white--text" depressed x-large color="#81D4FA"
                            @click="searchbtn"> <v-icon>mdi-magnify</v-icon> Search</v-btn>
                    </v-col>
                    <v-snackbar v-model="snackbar" :timeout=2000 centered color="red">
                        Incorrect Input, Please Check</v-snackbar>
                </v-row>
            </v-form>
        </v-card>

        <!-- <MainSaleReport v-show="tableflag"></MainSaleReport> -->
        <v-container v-show="tableflag"  class="mt-4">
            <v-card elevation="2" outlined>
                <v-card-title>
                    <v-spacer></v-spacer>
                    <v-btn class="deep-purple--text  elevation-10" @click="downloadCSV">Download</v-btn>
                </v-card-title>
                <v-text-field v-model="search" append-inner-icon="mdi-magnify" label="Search" outlined hide-details
                    rounded></v-text-field>
                <v-data-table :headers="headers" :items="arr" :search="search"></v-data-table>
                <v-snackbar v-model="snackbar1" :timeout=2000 centered color="orange  "> No Item In Sales Report</v-snackbar>
            </v-card>
        </v-container>
    </v-container>
</template>

<script>
import Papa from 'papaparse';
import EventServices from '../services/EventServices';

// import MainSaleReport from '../components/SalesReport/MainSaleReport.vue';
export default {
    data() {
        return {
            fromDate: null,
            toDate: null,
            fromDateMenu: false,
            toDateMenu: false,
            tableflag: false,
            search: '',
            rawData: [],
            snackbar: false,
            snackbar1: false,

        };
    },
    methods: {
        downloadCSV() {
            if (this.arr.length == 0) {
                this.snackbar1=true
                return
            }
            const csv = Papa.unparse(this.arr, {
                header: true
            });

            const blob = new Blob([csv], { type: 'text/csv;charset=utf-8;' });
            const link = document.createElement('a');
            const url = URL.createObjectURL(blob);
            link.setAttribute('href', url);
            link.setAttribute('download', 'data.csv');
            link.style.visibility = 'hidden';
            document.body.appendChild(link);
            link.click();
            document.body.removeChild(link);
        },
        searchbtn() {
            // console.log(this.fromDate)
            // console.log(this.toDate)
            if (this.fromDate !== null && this.toDate !== null && this.fromDate <= this.toDate) {

                EventServices.getSalesReport(this.fromDate, this.toDate)
                    .then(response => {
                        // console.log(response.data)
                        this.rawData = response.data.salesArr;
                        this.tableflag = true; // Assuming salesArr is the array in the response
                        // console.log(this.rawData)
                    })
                    .catch(error => {
                        console.error("API Error:", error);
                    });
                // this.$refs.form.reset();
            } else {
                this.snackbar = true;
            }
        },
        formatDate(date) {
            // Parse the ISO 8601 date string
            const parsedDate = new Date(date);

            // Extract the day, month, and year
            const day = String(parsedDate.getDate()).padStart(2, '0'); // Pad with leading zero if needed
            const month = String(parsedDate.getMonth() + 1).padStart(2, '0'); // Months are zero-based, so add 1 and pad with leading zero if needed
            const year = parsedDate.getFullYear();

            // Format the date as dd/mm/yyyy
            return `${day}/${month}/${year}`;
        },



    },
    components: {
        // MainSaleReport,
    },
    computed: {
        headers() {
            return [
                { text: 'Bill No', value: 'BillNo' },
                { text: 'Bill Date', value: 'BillDate' },
                { text: 'Medicine Name', value: 'MedicineName' },
                { text: 'Quantity', value: 'Quantity' },
                { text: 'Amount', value: 'Amount' },
            ]
        },

        arr() {
            if (!this.rawData) {
                return [];
            }else{
            return this.rawData.map(item => ({
                BillNo: item.billNo,
                BillDate: item.billDate,
                MedicineName: item.medicineName,
                Quantity: item.quantity,
                Amount: item.amount,
            }));
        }
        },


    },
};
</script>
