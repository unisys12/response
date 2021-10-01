export class EventBridge {
	private entity;
	private eventHandlers;

	constructor(entity, dispatch, events = []) {
		this.entity = entity;
		this.eventHandlers = [];

		if (events.length > 0) {
			const eventMap = {};

			events.forEach((event) => {
				if (!(event in eventMap)) {
					const handler = function (e) {
						dispatch(event, e);
					};

					this.eventHandlers.push({
						event: event,
						handler
					});

					entity.on(event, handler);
					eventMap[event] = handler;
				}
			});
		}
	}

	unregister() {
		this.eventHandlers.forEach((entry) => {
			this.entity.off(entry.event, entry.handler);
		});
	}
}
