describe('login-test', () => {

    const homeUrl = Cypress.env('homeUrl');
    const apiUrl = Cypress.env('apiUrl');

    before(() => {
        cy.request({
            method: 'POST',
            url: apiUrl + '/user',
            body: {
                "id": 0,
                "name": "test name",
                "last_name": "test last name",
                "dni": "123456",
                "email": "test@test.com",
                "password": "test"
            },
            failOnStatusCode: false // Ensure Cypress doesn't fail the test if the status code isn't 200
        })
    });

    beforeEach(() => {
        cy.visit(homeUrl)
    })

    it('should fail to login on user not registered', () => {

        /* ==== Generated with Cypress Studio ==== */
        cy.get('[href="/login"] > .boton').should('have.text', 'Iniciar sesion');
        cy.get('[href="/login"] > .boton').click();
        cy.get(':nth-child(1) > input').clear();
        cy.get(':nth-child(1) > input').type('test123456@test.com');
        cy.get(':nth-child(2) > input').clear();
        cy.get(':nth-child(2) > input').type('test123');
        cy.get('form > button').click();
        cy.get('.error-message').should('have.text', 'user not registered');
        /* ==== End Cypress Studio ==== */
    });

    it('should fail to login on wrong password', () => {

        /* ==== Generated with Cypress Studio ==== */
        cy.get('[href="/login"] > .boton').should('have.text', 'Iniciar sesion');
        cy.get('[href="/login"] > .boton').click();
        cy.get(':nth-child(1) > input').clear();
        cy.get(':nth-child(1) > input').type('test@test.com');
        cy.get(':nth-child(2) > input').clear();
        cy.get(':nth-child(2) > input').type('test123');
        cy.get('form > button').click();
        cy.get('.error-message').should('have.text', 'incorrect password');
        /* ==== End Cypress Studio ==== */
    });

    it('should succeed to login user', () => {

        /* ==== Generated with Cypress Studio ==== */
        cy.get('[href="/login"] > .boton').should('have.text', 'Iniciar sesion');
        cy.get('[href="/login"] > .boton').click();
        cy.get(':nth-child(1) > input').clear();
        cy.get(':nth-child(1) > input').type('test@test.com');
        cy.get(':nth-child(2) > input').clear();
        cy.get(':nth-child(2) > input').type('test');
        cy.get('form > button').click();
        cy.get('[href="/profile"] > .boton').should('have.text', 'Hola test name');
        /* ==== End Cypress Studio ==== */
    });

})