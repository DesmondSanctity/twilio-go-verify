const SMSVerification = {
 template: `
        <div class="max-w-md mx-auto bg-white rounded-xl shadow-md overflow-hidden md:max-w-2xl m-4">
            <div class="p-8">
                <h2 class="text-2xl font-bold mb-4">SMS Verification</h2>
                <form @submit.prevent="verifySMS">
                    <div class="mb-4">
                        <label class="block text-gray-700 text-sm font-bold mb-2" for="code">Enter SMS Code</label>
                        <input class="form-input" id="code" v-model="code" type="text" required>
                    </div>
                    <button class="btn" type="submit">Verify</button>
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
