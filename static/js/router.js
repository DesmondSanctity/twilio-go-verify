const router = new VueRouter({
 routes: [
  { path: '/', component: Signup },
  { path: '/login', component: Login },
  { path: '/sms-verification', component: SMSVerification },
  { path: '/authy-setup', component: AuthySetup },
  { path: '/dashboard', component: Dashboard },
 ],
});
