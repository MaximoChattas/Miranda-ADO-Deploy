describe('hotel-test', () => {

    const homeUrl = Cypress.env('homeUrl');
    const apiUrl = Cypress.env('apiUrl');
    let createdHotel;

    beforeEach(() => {
        cy.visit(homeUrl)
    });

    const loginAsAdmin = () => {
        const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmF0aW9uIjoxNzMyODUyNzIwLCJpZCI6MSwicm9sZSI6IkFkbWluIn0.LmivPdAyngWNEresG9T2a7tBiNdmj5IkJ-UOn03Ai4c"
        const user = {"id":1,"name":"Maximo","last_name":"Chattas","dni":"44347116","email":"maxichattas@gmail.com","role":"Admin"}

        window.localStorage.setItem('token', token);
        window.localStorage.setItem('userProfile', JSON.stringify(user));
    };


    it('should display hotel admin panel when logged in as admin user', () => {
        loginAsAdmin();

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
        loginAsAdmin();

        
        cy.get('[href="/profile"] > .boton').click();
        cy.get('.contenedorAdmin > :nth-child(2) > :nth-child(1)').click();
        cy.get('form > button').click();
        cy.get('.error-message').should('have.text', 'Complete todos los campos requeridos');
        
    })

    it('should succeed to create hotel', () => {
        loginAsAdmin();

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
            url: `${apiUrl}/hotel/${createdHotel.id}`,
        }).then((response) => {
            expect(response.status).to.eq(200); // Confirm deletion was successful
        });
    });
})