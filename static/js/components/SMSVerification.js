const SMSVerification = {
 template: `
        <div class="max-w-md mx-auto bg-white rounded-xl shadow-md overflow-hidden md:max-w-2xl m-4">
            <div class="p-8">
                <h2 class="text-2xl font-bold mb-4">SMS Verification</h2>
                <form @submit.prevent="verifySMS">
                    <div class="mb-4">
                        <label class="block text-gray-700 text-sm font-bold mb-2" for="code">Enter SMS Code</label>
                        <input class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                        id="code" v-model="code" type="text" required>
                    </div>
                    <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                    type="submit">Verify</button>
                </form>
            </div>
        </div>
    `,
 data() {
  return {
   code: '',
   user: JSON.parse(localStorage.getItem('user')),
  };
 },
 async mounted() {
  if (!this.user) {
   alert('Login or signup first!');
   this.$router.push('/');
   return;
  }
 },
 methods: {
  async verifySMS() {
   try {
    const user = JSON.parse(localStorage.getItem('user'));
    await axios.post('/api/verify/verify-sms', {
     email: user.email,
     code: this.code,
    });
    this.user.smsEnabled = true;
    localStorage.setItem('user', JSON.stringify(this.user));
    this.$router.push('/authy-setup');
   } catch (error) {
    alert('SMS verification failed: ' + error);
   }
  },
 },
};
