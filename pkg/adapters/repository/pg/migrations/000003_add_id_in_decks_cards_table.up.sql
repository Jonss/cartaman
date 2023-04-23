ALTER TABLE decks_cards ADD COLUMN id bigserial NOT NULL;
ALTER TABLE decks_cards ADD CONSTRAINT decks_cards_id_pk primary key (id);