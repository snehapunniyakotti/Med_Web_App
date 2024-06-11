<template>
  <v-container elevation-2 class="mt-16 white">
    <!-- <v-card> -->
    <v-row>
      <v-col cols="12">
        <h1>Stock Entry</h1>
        <v-spacer></v-spacer>
        <v-container>
          <v-row>
            <v-spacer></v-spacer>
            <v-btn @click="openDialog" class="white--text" color="#81D4FA">Add <v-icon>mdi-pill-multiple</v-icon>
            </v-btn>
          </v-row>
          <v-dialog v-model="dialog" max-width="600px">
            <v-card>
              <v-card-title class="light-blue lighten-3 white--text">
                <span class="headline">Add Stock</span>
              </v-card-title>
              <v-card-text class="mt-6">
                <v-text-field v-model.lazy="medicineName" outlined label="Medicine Name"></v-text-field>
                <v-text-field v-model.lazy="brand" outlined label="Brand"></v-text-field>
              </v-card-text>
              <v-card-actions>
                <v-btn class="white--text" color="#81D4FA"
                  @click="closeDialog">Close<v-icon>mdi-close-circle</v-icon></v-btn>
                <v-btn class="white--text" color="#81D4FA" @click="addStocks">Add<v-icon>mdi-pill-multiple</v-icon>
                </v-btn>
                <!-- <span class="text-h4 text-uppercase error--text ml-10 font-weight-bold">{{ msg }}</span> -->
                <v-snackbar v-model="snackbar1" :timeout="timeout" centered color="green">{{ msg }}</v-snackbar>
                <v-snackbar v-model="snackbar4" :timeout="timeout" centered color="red">{{ msg }}</v-snackbar>
                <v-snackbar v-model="snackbar5" :timeout="timeout" centered color="red">Incorrect Input, Please
                  Check</v-snackbar>
              </v-card-actions>
            </v-card>
          </v-dialog>
        </v-container>
      </v-col>
    </v-row>

    <v-row>
      <v-col cols="3">
        <v-select outlined v-model="selectedItem" :items="medicines" label="Medicine Name"
          @change="fetchMedicineDetails"></v-select>
      </v-col>
      <v-col cols="3">
        <v-text-field outlined v-model="selectedItemBrand" label="Brand" readonly></v-text-field>
      </v-col>
      <!-- </v-row>
  
      <v-row> -->
      <v-col cols="3">
        <v-text-field outlined v-model.number="quantity" label="Quantity" type="number"></v-text-field>
      </v-col>
      <v-col cols="3">
        <v-text-field outlined v-model.number="price" label="Unit Price" type="number"></v-text-field>
      </v-col>
    </v-row>

    <v-row class="">
      <v-col cols="12" class="">
        <v-btn block class="white--text" color="#81D4FA" @click="update">Update</v-btn>
      </v-col>
    </v-row>
    <v-row>
      <!-- <v-spacer></v-spacer> -->
      <!-- <span class="text-h4 text-uppercase success--text   font-weight-black ">{{ msgupdate }}</span> -->
      <v-snackbar v-model="snackbar2" :timeout="timeout" centered color="green">{{
        msgupdate
      }}</v-snackbar>
      <v-snackbar v-model="snackbar3" :timeout="timeout" centered color="red">Incorrect Input, Please Check</v-snackbar>
    </v-row>
    <!-- </v-card> -->
  </v-container>
</template>

<script>
import EventServices from '../services/EventServices';
export default {
  data() {
    return {
      // Sample array of items
      selectedItem: "",
      selectedItemBrand: "",
      quantity: null,
      price: null,
      dialog: false,
      medicineName: "",
      brand: "",
      msg: "",
      msgupdate: "",
      timeout: 2000,
      snackbar1: false,
      snackbar2: false,
      snackbar3: false,
      snackbar4: false,
      snackbar5: false,
      medicines: [],
    };
  },
  methods: {
    openDialog() {
      this.msg = "";
      this.medicineName = "";
      this.brand = "";
      this.dialog = true;
    },
    closeDialog() {
      this.dialog = false;
    },
    checkinguserid(v) {
      return !!v && /^[A-Za-z]{2}/.test(v);
    },
    checkingpass(v) {
      return !!v && /^[A-Za-z]/.test(v);
    },

    addStocks() {
      this.medicineName = this.medicineName.trim();
      this.brand = this.brand.trim();
      if (
        !this.medicineName ||
        !this.brand ||
        !this.checkinguserid(this.medicineName) ||
        !this.checkingpass(this.brand)
      ) {
        this.snackbar5 = true;
        this.medicineName = "";
        this.brand = "";
        return;
      }
      // console.log(this.medicineName + 'dgf' + this.brand);
      EventServices.addStock(this.medicineName, this.brand)
        .then(response => {
          const data = response.data;
          if (data.resp.status === "S") {
            this.msg = response.data.msg;
            this.snackbar1 = true;
            this.medicineName = "";
            this.brand = "";
          } else {
            this.msg = response.data.resp.errMsg;
            this.snackbar4 = true
          }

        })
        .catch(error => {
          console.error("API Error :" + error);
          this.msg = "Error adding stock";
          this.snackbar1 = true;
        });
    },
    fetchMedicines() {
      EventServices.getMedicineNames()
        .then(response => {
          const data = response.data;
          // console.log('API Response:', data); // Debug log
          if (data.resp.status === "S") {
            this.medicines = data.medArr.map(med => med.medName); // Ensure correct mapping
            // console.log('Medicines:', this.medicines); // Debug log
          } else {
            console.error('Error:', data.resp.errMsg);
          }
        })
        .catch(error => {
          console.error('API Error:', error);
        });
    },
    fetchMedicineDetails() {
      EventServices.getMedicineDetails(this.selectedItem)
        .then(response => {
          const data = response.data;
          if (data.resp.status === "S") {
            this.selectedItemBrand = data.brand;
            this.price = data.unitPrice;
          } else {
            console.error('Error fetching medicine details:', data.MResp.ErrMsg);
          }
        })
        .catch(error => {
          console.error('API Error:', error);
        });
    },
    update() {
      if (this.price <= 0 || this.quantity <= 0) {
        // Show error message if quantity or price is invalid
        this.snackbar3 = true;
        return;
      }

      // Prepare data object to send to API
      let obj1 = {
        medicineName: this.selectedItem,
        quantity: this.quantity,
        unitPrice: this.price,
        brand: this.selectedItemBrand,
      };
      //  let jsondata = JSON.stringify(obj1)
      // Call the API using EventServices
      EventServices.updateStock(obj1)
        .then(response => {
          // Handle API response
          const data = response.data;
          if (data.resp.status === "S") {
            // Show success message
            this.msgupdate = "Stock Updated";
            this.snackbar2 = true;
          } else {
            // Show error message if API response indicates an error
            console.error('Error:', data.resp.errMsg);
            this.snackbar3 = true;
          }
        })
        .catch(error => {
          // Handle API error
          console.error('API Error:', error);
          this.snackbar3 = true;
        });

      // Reset form inputs after API call
      this.selectedItem = "";
      this.selectedItemBrand = "";
      this.quantity = null;
      this.price = null;
    },
  },
  mounted() {
    this.fetchMedicines(); // Fetch the list of medicines when component is mounted
  }
};
</script>
