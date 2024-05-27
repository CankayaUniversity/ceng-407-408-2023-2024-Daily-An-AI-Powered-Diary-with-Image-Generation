describe('Login Test', () => {
  beforeAll(async () => {
    await device.launchApp();
  });

  afterAll(async () => {
    await device.terminateApp();
  });

  it('should login successfully with valid credentials', async () => {
    // Step A01: Go to the login page
    // Assuming the app launches directly to the login page

    // Step A02: Enter a valid user id
    await element(by.id('loginUserId')).tap();
    await element(by.id('loginUserId')).typeText('valid_user_id');
    await device.pressBack(); // Dismiss the keyboard, if necessary

    // Step A03: Enter the valid password for this user
    await element(by.id('loginPassword')).tap();
    await element(by.id('loginPassword')).typeText('valid_password');
    await device.pressBack(); // Dismiss the keyboard, if necessary

    // Step A04: Click on the “Login” button
    await element(by.id('loginButton')).tap();

    // Step V01: Observe that the login is successful and the user dashboard appears
    await expect(element(by.id('userDashboard'))).toBeVisible();
  });

  afterEach(async () => {
    // Cleanup: Logout
    await element(by.id('logoutButton')).tap();
    await expect(element(by.id('loginButton'))).toBeVisible(); // Ensure we are back to the login screen
  });
});

