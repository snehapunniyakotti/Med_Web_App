import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    Login: [
      {
        User_Id: "M",
        Password: 1,
        Role: 'Manager'
      },
      {
        User_Id: "B",
        Password: 2,
        Role: 'Biller'
      },
      {
        User_Id: "S",
        Password: 3,
        Role: 'SystemAdmin'
      },
      {
        User_Id: "I",
        Password: 4,
        Role: 'Inventry'
      },
    ],
    Login_History: [
      {
        User_Id: "S",
        Type: 'Login',
        Date: new Date(),
      },
      {
        User_Id: "B",
        Type: 'Login',
        Date: new Date(),
      },
      {
        User_Id: "M",
        Type: 'Logout',
        Date: new Date(),
      },
    ],
    Stock: [
      {
        MedicineName: "Aspirin",
        Quantity: 145,
        Unit_Price: 2,
        Brand: "Ecosprin"
      },
      {
        MedicineName: "Ibuprofen",
        Quantity: 40,
        Unit_Price: 5,
        Brand: "Brufen"
      },
      {
        MedicineName: "Ranitidine",
        Quantity: 15,
        Unit_Price: 3,
        Brand: "Zinetac"
      },
      {
        MedicineName: "Diclofenac",
        Quantity: 100,
        Unit_Price: 2,
        Brand: "Voveran"
      }
    ],
    Medicine_Master: [
      
      {
        MedicineName: "Ibuprofen",
        Brand: "Brufen"
      },
      {
        MedicineName: "Aspirin",
        Brand: "Ecosprin"
      },
      {
        MedicineName: "Ranitidine",
        Brand: "Zinetac"
      },
      {
        MedicineName: "Diclofenac",
        Brand: "Voveran"
      }
    ],
    Bill_Master: [
      {
        Bill_No: "B001",
        Bill_Date: "11/03/2024",
        Bill_Amount: 200,
        Bill_Gst: 8.9,
        Net_Price: 98.2,
        User_Id: "B"
      },
      {
        Bill_No: "B002",
        Bill_Date: "07/04/2024",
        Bill_Amount: 150,
        Bill_Gst: 29,
        Net_Price: 160,
        User_Id: "B"
      },
    ],
    Bill_Details: [
      {
        Bill_No: 'B001',
        MedicineName: 'Ibuprofen',
        Quantity: 30,
        Unit_Price: 4,
        Amount: 120
      },
      {
        Bill_No: 'B002',
        MedicineName: 'Aspirin',
        Quantity: 60,
        Unit_Price: 3.5,
        Amount: 210
      }
    ],
  },
  mutations: {
    login_entry(state, obj) {
      // console.log("mutated")
      state.Login_History.push(obj);
    },
    addUser(state, obj) {
      // console.log(obj)
      state.Login.push(obj);
    },
    updateStock(state, obj) {
      let arr = state.Stock;
      for (let i = 0; i < arr.length; i++) {
        if (arr[i].MedicineName == obj.MedicineName) {
          arr[i].Quantity = parseInt(arr[i].Quantity) + parseInt(obj.Quantity);
          arr[i].Unit_Price = parseFloat(obj.Unit_Price);

          // console.log(state.Stock[i]);
          return
        }
      }

      let obj1 = {
        MedicineName: obj.MedicineName,
        Quantity: parseInt(obj.Quantity),
        Unit_Price: parseFloat(obj.Unit_Price),
        Brand: obj.Brand
      };
      // console.log('hi')
      // console.log(obj1)

      state.Stock.push(obj1);
    },
    updateMM(state, obj) {
      state.Medicine_Master.push(obj);
    },
    updateBM(state, obj) {
      state.Bill_Master.push(obj);
    },
    updateBD(state, obj) {
      state.Bill_Details.push(obj);


      let stock = state.Stock.find(med => med.MedicineName == obj.MedicineName);
      stock.Quantity -= obj.Quantity;
    }
  },
  actions: {},
  modules: {},
});
