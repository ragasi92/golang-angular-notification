export class AppNotification {
    constructor(
        public id: string,
        public user?: string,
        public action?: string,
        public creationDate?: Date,
        public status?: string,
    ) { }
}

