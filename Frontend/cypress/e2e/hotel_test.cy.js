describe('hotel-test', () => {

    const homeUrl = Cypress.env('homeUrl');
    const apiUrl = Cypress.env('apiUrl');
    let createdHotel;

    beforeEach(() => {
        cy.visit(homeUrl)
    });

    it('should display hotel admin panel when logged in as admin user', () => {
        // Login as admin user
        cy.get('[href="/login"] > .boton').click();
        cy.get(':nth-child(1) > input').clear();
        cy.get(':nth-child(1) > input').type('maxichattas@gmail.com');
        cy.get(':nth-child(2) > input').clear();
        cy.get(':nth-child(2) > input').type('admin');
        cy.get('form > button').click();

        /* ==== Generated with Cypress Studio ==== */
        cy.get(':nth-child(1) > .card > .card-body > .card-title > a').click();
        cy.get('.description > :nth-child(7)').should('be.visible');
        /* ==== End Cypress Studio ==== */
    });

    it('should not display hotel admin panel when not logged in as admin user', () => {

        /* ==== Generated with Cypress Studio ==== */
        cy.get(':nth-child(1) > .card > .card-body > .card-title > a').click();
        cy.get('.description > :nth-child(7)').should('not.exist');
        /* ==== End Cypress Studio ==== */
    });

    it('should fail to create hotel with empty fields', () => {
        // Login as admin user
        cy.get('[href="/login"] > .boton').click();
        cy.get(':nth-child(1) > input').clear();
        cy.get(':nth-child(1) > input').type('maxichattas@gmail.com');
        cy.get(':nth-child(2) > input').clear();
        cy.get(':nth-child(2) > input').type('admin');
        cy.get('form > button').click();
        /* ==== Generated with Cypress Studio ==== */
        cy.get('[href="/profile"] > .boton').click();
        cy.get('.contenedorAdmin > :nth-child(2) > :nth-child(1)').click();
        cy.get('form > button').click();
        cy.get('.error-message').should('have.text', 'Complete todos los campos requeridos');
        /* ==== End Cypress Studio ==== */
    })

    it('should succeed to create hotel', () => {
        // Login as admin user
        cy.get('[href="/login"] > .boton').click();
        cy.get(':nth-child(1) > input').clear();
        cy.get(':nth-child(1) > input').type('maxichattas@gmail.com');
        cy.get(':nth-child(2) > input').clear();
        cy.get(':nth-child(2) > input').type('admin');
        cy.get('form > button').click();

        cy.get('[href="/profile"] > .boton').click();
        cy.get('.contenedorAdmin > :nth-child(2) > :nth-child(1)').click();

        cy.get(':nth-child(1) > input').clear();
        cy.get(':nth-child(1) > input').type('Test Hotel');

        cy.get('form > :nth-child(2) > input').clear();
        cy.get('form > :nth-child(2) > input').type('Test Hotel St');

        cy.get('form > :nth-child(3) > input').clear();
        cy.get('form > :nth-child(3) > input').type('123');

        cy.get('form > :nth-child(4) > input').clear();
        cy.get('form > :nth-child(4) > input').type('23');

        cy.get('form > :nth-child(5) > input').clear();
        cy.get('form > :nth-child(5) > input').type('45000');

        cy.get('textarea').click();

        // Set up intercept for hotel creation before submitting
        cy.intercept('POST', apiUrl+'/hotel').as('createHotel');

        cy.get('form > button').click();

        // Wait for the request and capture the response
        cy.wait('@createHotel').then((interception) => {
            createdHotel = interception.response.body;
            expect(interception.response.statusCode).to.eq(201); // Check if creation was successful
            cy.wrap(createdHotel).as('createdHotel');
        });

        cy.intercept('GET', homeUrl, (req) => {
            req.on('response', (res) => {
                // Ensure response status is 200
                expect(res.statusCode).to.eq(200);
            });
        }).as('getHomepage');
        cy.visit(homeUrl);
        cy.wait('@getHomepage');

        cy.get('.row > .col-md-4:last')
            .find('.card > .card-body > .card-title > a')
            .should('have.text', 'Test Hotel');

        cy.get('.row > .col-md-4:last')
            .find('.card > .card-body > :nth-child(2)')
            .should('have.text', 'Dirección: Test Hotel St 123');

        cy.get('.row > .col-md-4:last')
            .find('.card > .card-body > :nth-child(3)')
            .should('have.text', '$45000');
    });

    after(() => {
        cy.request({
            method: 'DELETE',
            url: `${apiUrl}/hotel/${createdHotel.id}`, // Replace with the correct URL structure
        }).then((response) => {
            expect(response.status).to.eq(200); // Confirm deletion was successful
        });
    });
})