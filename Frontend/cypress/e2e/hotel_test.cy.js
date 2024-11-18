Cypress.Commands.add('loginAsAdmin', () => {
    cy.request({
        method: 'POST',
        url: Cypress.env('apiUrl') + '/login',
        url: apiUrl + '/login',
        body: {
            email: 'maxichattas@gmail.com',
            password: 'admin'
        }
    }).then((response) => {
        if (response.status === 202) {
            const { token, user } = response.body;
            cy.log(JSON.stringify(response));
            window.localStorage.setItem('token', token);
            window.localStorage.setItem('userProfile', JSON.stringify(user));
        } else {
            throw new Error('Login failed during test setup');
        }
    });
});

describe('hotel-test', () => {

    const homeUrl = Cypress.env('homeUrl');
    const apiUrl = Cypress.env('apiUrl');
    let createdHotel;

    beforeEach(() => {
        cy.visit(homeUrl)
    });

    it('should display hotel admin panel when logged in as admin user', () => {

        cy.loginAsAdmin();

        cy.intercept('GET', apiUrl+'/hotel/*', (req) => {
            req.on('response', (res) => {
                // Ensure response status is 200
                expect(res.statusCode).to.eq(200);
            });
        }).as('getHotel');
        cy.get(':nth-child(1) > .card > .card-body > .card-title > a').click();
        cy.wait('@getHotel');

        cy.get('.description > :nth-child(7)').should('be.visible');
        
    });

    it('should not display hotel admin panel when not logged in as admin user', () => {
        cy.intercept('GET', apiUrl+'/hotel/*', (req) => {
            req.on('response', (res) => {
                expect(res.statusCode).to.eq(200);
            });
        }).as('getHotel');
        cy.get(':nth-child(1) > .card > .card-body > .card-title > a').click();
        cy.wait('@getHotel');

        cy.get('.description > :nth-child(7)').should('not.exist');
        
    });

    it('should fail to create hotel with empty fields', () => {
        cy.loginAsAdmin();
        
        cy.get('[href="/profile"] > .boton').click();
        cy.get('.contenedorAdmin > :nth-child(2) > :nth-child(1)').click();
        cy.get('form > button').click();
        cy.get('.error-message').should('have.text', 'Complete todos los campos requeridos');
        
    })

    it('should succeed to create hotel', () => {
        cy.loginAsAdmin();

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

        cy.intercept('POST', apiUrl+'/hotel').as('createHotel');
        cy.get('form > button').click();
        cy.wait('@createHotel').then((interception) => {
            createdHotel = interception.response.body;
            expect(interception.response.statusCode).to.eq(201);
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