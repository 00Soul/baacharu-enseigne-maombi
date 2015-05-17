(function() {
    var injectParams = ['$scope'];

    var BoardController = function($scope) {
        var user = {
            id: "jns3ajfmi5r6vnzaskfoepj5h2sf7a",
            profile: {
                email: "jacob7227@aol.com",
                username: "jacobp",
                displayname: "Jacob"
            }
        };

        $scope.board = {
            title: "Sample Board",
            ownedby: user.id,
            createdby: user.id,
            columns: [
                {
                    title: "To Do",
                    wiplimit: 5,
                    cards: [
                        {
                            createdby: user.id,
                            createdwhen: "2015-05-02T10:53:06-04:00"
                        }, {
                            createdby: user.id,
                            createdwhen: "2015-05-01T13:21:34-04:00"
                        }
                    ]
                }, {
                    title: "Doing",
                    wiplimit: 3,
                    cards: [
                        {
                            createdby: user.id,
                            createdwhen: "2015-05-01T13:04:57-04:00"
                        }
                    ]
                }, {
                    title: "Done",
                }
            ],
            cards: [
                {
                    createdby: user.id,
                    createdwhen: "2015-05-02T10:53:06-04:00",
                    stage: 0
                }, {
                    createdby: user.id,
                    createdwhen: "2015-05-01T13:21:34-04:00",
                    stage: 0
                }, {
                    createdby: user.id,
                    createdwhen: "2015-05-01T13:04:57-04:00",
                    stage: 1
                }
            ]
        };
    };

    BoardController.$inject = injectParams;

    angular.module('winch').controller('BoardController', BoardController);
}());
