# Nuky Development Design Sheet

### Scheduled Jobs

1. nukeState() / 24h 12:00 UTC:

### Functional Requirements

1. auth:
   login()
   logout()
   register()

2. move:
   both functions below requires can_attack true
   every move should be checked by recaptcha
   attack()
   support()

3. state:
   getStates()

4. game:
   getGame()

### Non Functional Requiremets

1. latency: 200ms

2. rps: 1/2 x 5 x 100 = 250 RPS

3. memory: 100mb

4. cpu: graviton-small 60%

### Entities

1. User
   email: string, UNIQUE
   password: hashed string
   last_move_date: Date

2. State
   id: int, UNIQUE
   state_name: string
   state_color_hex: string
   attack_count: int
   support_count: int

3. Game
   nuked_states: State[]
