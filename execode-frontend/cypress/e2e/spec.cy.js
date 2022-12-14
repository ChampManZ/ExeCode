describe('My First Test', () => {
  it('Visits the Landing Page', () => {
    cy.visit('http://localhost:3001/')
    
    cy.contains('Welcome to ExeCode!')

    cy.contains('Login / Sign-Up').click()
  })
})

describe('My Second Test', () => {
  it('Visits Homepage', () => {
    cy.visit('http://localhost:3001/home')

    cy.contains('To Do...')

    cy.contains('Announcement')
  })
})

describe('My Third Test', () => {
  it('Visits Course', () => {
    cy.visit('http://localhost:3001/courses')

    cy.contains('Courses')
  })
})

describe('My Third Test', () => {
  it('Visits Course Enroll', () => {
    cy.visit('http://localhost:3001/courses/1')

    cy.contains('Fundamental of Programming')

    cy.contains('This course introduces basic concepts of computer programming such as elementary programming, data types, expressions, simple algorithms and problem solving involving sequential statements, conditionals and iterations. Students learn routines or methods as fundamental concepts and practice using strings, arrays, lists, maps or dictionaries, pre-defined libraries and classes, abstraction mechanisms and basic object- oriented programming concepts. Students will practice related activities of software development life cycle such as system requirement analysis, debugging, testing and validation.')

    cy.contains('Modules').click()

    cy.url().should('include', '/courses/1/module')

    cy.contains('Modules')

    cy.contains('Module 1').click()
  })
})

describe('My Fourth Test', () => {
  it('Visits Code Runner', () => {
    cy.visit('http://localhost:3001/coderunplayground')

    cy.contains('Language:')

    cy.contains('Untitled Problem Statement')

    cy.contains('Test Case')

    cy.contains('Output')

    cy.contains('To add test cases, please use ";" to separate between cases.')

    cy.contains('Example: [2, 5];[5,-10]')

    cy.contains('Input n = 0')

    cy.contains('Add Test Case').click()

    cy.get('.editor-test').type('\n\n5;15')

    cy.contains('Add').click()

    cy.get('.editor-main').type('console.log(process.argv)')

    cy.get('.language-select').click()

    cy.contains('JavaScript').click()
  })
})