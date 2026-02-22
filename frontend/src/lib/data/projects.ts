export interface ProjectEntry {
	id: string;
	name: string;
	shortDescription: string;
	longDescription: string;
	images?: string[];
}

export const projects: ProjectEntry[] = [
	{
		id: 'penncloud',
		name: 'PennCloud',
		shortDescription: 'A highly available distributed system for file storage, replication and recovery using C++',
		longDescription: `For the CIS 5050: Software Systems project, Me and my team developed a google mail and drive like service using just C++ and gRPC. This distributed system could store files in chunks and access them using the basic GET, PUT, DELETE and CPUT operations. The system could load balance requests across its many frontend and backend servers. We also added functionaliy to use uploaded files as context to ask questions using an LLM wrapper and a google docs like system that allowed for collaborative editing access.
		
		My primary responsibilities included:
		- Designing and implementing a bigtable-like key-value store that could be used to store files (in chunks), emails and other documents of all types.
		- Add Paxos based consenus model to ensure data consistency across all servers
		- Make sure the system was highly avaliable primarily for read operations
		- Uses the consensus model to also build a checkpointing based recovery mechanism
		
		Challenges:
		- Migrating to gRPC mid-project caused a lot of pain and required a lot of debugging and refactoring of the existing codebase (along with the added confusion of using a new protocol ofc)
		- As we decided to have R=1, it meant that our Paxos consensus was heavily reliant on all server groups consistently having the most updated state, causing some lag`,
		images: []
	},
	{
		id: 'portfolio',
		name: 'portfolio-websites',
		shortDescription: 'This very website (portolio) to display personal information',
		longDescription: `Simple as it looks, this website has gone over 3 complete overhaul iterations over the last 4 years. Those are:
		
		Iteration 1: A simple React.js based, minimally designed portfolio application to display some information and make API calls to TMDB.
		Iteration 2: Migrated to a next.js based, more feature rich application. Better overall design, Firebase authentication and firestore based shared space for blog writing
		Iteration 3 (Current): Migrated to a SvelteKit based, more feature rich application. Better overall design, improved UI/UX, and a new stack (SvelteKit, Tailwind, Supabase, Vercel).`},
	{
		id: 'herehackathon',
		name: 'herehackathon',
		shortDescription: 'A navigation system built for the Here Hackathon @ IIT Bombay. Powered by react.js, tailwind and Here Maps API',
		longDescription: `For the Here Hackathon @ IIT Bombay, Me and my team built a navigation system that allowed for the navigation of a vehicle. We used react.js, tailwind and Here Maps API to build the frontend and backend respectively.

		My primary responsibilities included:
		- Team coordination
		- Frontend development of components
		- and linking the middleware API to the frontend

		Challenges:
		- The Here Maps API was not very well documented and had a lot of bugs, so we had to spend a lot of time debugging and finding workarounds
		- Since we were only a two-person team competing against 4+ member teams, we had to be quite selective with the functionalities we decided to implement. As a result, we had to make significant compromises that possibly docked our effective total.
		`,
	},
	{
		id: 'fleetwise',
		name: 'fleetwise',
		shortDescription: 'A fleet management system built for the Afourathon 3.0. Powered by Next.js and FastAPI',
		longDescription: `For Afourathon 3.0 online, Me and my team built a fleet management system that allowed for the management of a fleet of vehicles. We used Next.js and FastAPI to build the frontend and backend respectively. We also used the OpenAI API to build the AI powered features of the system.
		
		My primary responsibilities included:
		- Team management and planning/coordination
		- Building the backend using FastAPI
		- Building the frontend components for cabs and driver cards sections
		`,
	}
];
