<template>
  <v-container class="pt-0 mt-10">

    <v-layout>
      <v-flex lg12 class="d-flex justify-end">
      </v-flex>
    </v-layout>
    <v-card class=" pa-2 mt-3" rounded="0">
      <v-layout>
        <v-flex md-2>
          <preview :Previewarr="userArr" :Total="tp" :GST="gst" :NetPrice="netpayable" class="mt-2" />
        </v-flex>
        <v-flex md-2>
          <v-btn class="mr-6 mb-5  white--text mt-2 ml-3 " color="#81D4FA" @click="saveBill">Save</v-btn>
        </v-flex>
        <v-flex md-8 class="mt-3">
          <span class="mt-2 mr-6 "><b>BILL NO : </b>{{ billCounter }}</span>
          <span class="mt-2 mr-6"><b>DATE : </b>{{ date }}</span>
          <span class="mt-2 mr-6"><b>TOTAL :</b>{{ tp }}</span>
          <span class="mt-2 mr-6"><b>GST (18%) :</b>{{ ' ' + (Math.trunc((tp * (18 / 100)) * 100)) / 100 + '' }}</span>
          <span class="mt-2 mr-6"><b>NetPayable : </b>{{ this.tp + Math.floor(tp * (18 / 100)) }}</span>

        </v-flex>
      </v-layout>
    </v-card>

    <v-card elevation="0" rounded="0">
      <v-row class="">
        <!-- <span>hi</span> -->
        <v-spacer></v-spacer>
        <v-btn class="white--text mr-16" color="#81D4FA" elevation="1" @click="downloadCSV">download</v-btn>
      </v-row>
      <v-card-title>
        <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details
          dense></v-text-field>
      </v-card-title>
      <v-data-table :headers="headers" :items="userArr" :search="search" dense rounded class=""></v-data-table>
    </v-card>

    <v-snackbar v-model="snackbar1" :timeout=2000 centered color="red"> Bill Not Saved select medicine</v-snackbar>
    <v-snackbar v-model="snackbar2" :timeout=2000 centered color="green"> Bill Saved</v-snackbar>
    <v-snackbar v-model="snackbar3" :timeout=2000 centered color="orange  "> No Item In Bill</v-snackbar>


    <!-- {{ this.$store.state.BillMaster }} -->

    <!-- {{'gst'+  Math.trunc(17.3245224*100)/100  }} -->

  </v-container>
</template>


<script>
import Papa from 'papaparse';
import preview from './preview.vue';
import EventServices from '../../../services/EventServices';

export default {
  name: "billtable",
  data() {
    return {
      search: '',
      headers: [
        { text: 'Medicine Name', value: 'medicineName', class: "light-blue lighten-3 white--text " },
        { text: 'Brand', value: 'brand', class: "light-blue lighten-3 white--text" },
        { text: 'Quantity', value: 'quantity', class: "light-blue lighten-3 white--text" },
        { text: 'Amount', value: 'amount', class: "light-blue lighten-3 white--text " },
      ],
      desserts: [],
      billno: 0,
      date: new Date().toJSON().slice(0, 10),
      total: 0,
      gst: 0,
      netpayable: 0,
      // gst: Math.floor(this.tp * (18 / 100)),
      // netpayable: this.tp + Math.floor(this.tp * (18 / 100)),
      // stockarr: this.$store.state.Stock,
      // medmaster: this.$store.state.MedicineMaster,
      // billMaster: this.$store.state.BillMaster,
      // billDetails: this.$store.state.BillDetails,
      // LoginHistory: this.$store.state.LoginHistory,
      UserId: '',
      obj: {
        snackbar: false,
        color: 'red',
        topvar: false,
        bottomvar: false,
        text: '',

      },
      // previewpop:[],
      billCounter: 1000,
      snackbar1: false,
      snackbar2: false,
      snackbar3: false,

    }
  },
  components: {
    preview,
  },
  props: {
    // billNumber: String,
    userArr: Array,
    tp: Number,

  },

  methods: {
    downloadCSV() {
      if (this.userArr.length == 0) {
        this.snackbar3=true
        return
      }
      const csv = Papa.unparse(this.userArr, {
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
    getGst() {
      // this.gst = Math.floor(this.tp * (18 / 100));
      this.gst = Math.trunc((this.tp * (18 / 100)) * 100) / 100
    },
    getNetPay() {
      this.netpayable = this.tp + Math.floor(this.tp * (18 / 100));
    },
    billdataEmpty() {
      this.gst = 0,
        this.netpayable = 0,
        this.$emit('emptypreview');
    },
    saveBill() {
      if (this.userArr.length == 0) {
        this.snackbar1=true
        return;
      }

      // console.log(this.billCounter);
      this.billCounter += 1;
      
      let billData = {
        // billNo: this.billNumber,
        
        billNo: this.billCounter,
        billAmount: this.tp,
        billGst: this.gst,
        netPrice: this.netpayable,
        userId: this.$route.params.Id,
        medArr: this.userArr.map(item => ({
          medicineName: item.medicineName,
          quantity: parseInt(item.quantity),
          unitPrice: parseFloat(item.unitPrice),
          amount: parseFloat(item.amount),
        }))
      };
      // console.log("userId in billtable "+billData.userId)
      // console.log(this.billCounter);
      
      // console.log(this.$route.params.Id);
      
      EventServices.saveBill(billData)
      .then(response => {
        if (response.data.resp.status === 'S') {
          this.$emit('billclosed');
            this.snackbar2=true;
          } else {
            console.error(response.data.resp.errMsg);
          }
        })
        .catch(error => {
          console.error("API error :", error);
        });
      },
  },
  watch: {
    tp() {
      this.gst = Math.floor(this.tp * (18 / 100));
      this.netpayable = this.tp + Math.floor(this.tp * (18 / 100));
    },
    
  }
}

</script>
