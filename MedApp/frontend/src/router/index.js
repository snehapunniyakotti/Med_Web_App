import Vue from "vue";
import VueRouter from "vue-router";
import LoginPage from "../views/LoginPage.vue";
import Dashboard from "../views/Dashboard.vue";
import Errormsg from "../views/ErrorMsg.vue";
import AddUser from "../views/AddUser.vue";
import stockview from "../views/StockView.vue"
import LoginHistory from "../views/LoginHistory.vue"
import SalesReport from "../views/SalesReport.vue";
import StockEntry from "../views/StockEntry.vue";
import BillEntry from "../views/BillEntry.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Login",
    component: LoginPage,
  },
  {
    path: "/dashboard",
    name: "dashboard",
    component: Dashboard ,
  },
  {
    path: "/Errormsg",
    name: "Errormsg",
    component: Errormsg ,
  },
  {
    path: "/adduser",
    name: "adduser",
    component: AddUser ,
  },
  {
    path: "/stockview",
    name: "stockview",
    component: stockview ,
  },
  {
    path: "/loginhistory",
    name: "loginhistory",
    component: LoginHistory ,
  },
  {
    path: "/salesreport",
    name: "salesreport",
    component: SalesReport ,
  },
  {
    path: "/stockentry",
    name: "stockentry",
    component: StockEntry ,
  },
  {
    path: "/billentry",
    name: "billentry",
    component: BillEntry ,
  },
 
];

const router = new VueRouter({
  routes,
});

export default router;
