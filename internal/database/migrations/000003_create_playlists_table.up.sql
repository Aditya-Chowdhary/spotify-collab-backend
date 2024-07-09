CREATE TABLE public.playlists (
	event_uuid uuid NOT NULL,
	playlist_id text NOT NULL,
	"name" citext NOT NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	updated_at timestamptz DEFAULT now() NOT NULL,
	CONSTRAINT playlists_pk PRIMARY KEY (playlist_id),
	CONSTRAINT playlists_unique UNIQUE (event_uuid, name),
	CONSTRAINT playlists_events_fk FOREIGN KEY (event_uuid) REFERENCES public.events(event_uuid) ON DELETE CASCADE ON UPDATE CASCADE
);