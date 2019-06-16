export default {
  albumData: [
    {
      id: 1,
      name: "Jakarta - Bandung",
      IsPaidOff: false
    },
    {
      id: 2,
      name: "Jakarta - Yogyakarta",
      IsPaidOff: true
    },
    {
      id: 3,
      name: "Liburan keluarga",
      IsPaidOff: true
    }
  ],
  albumDetail: {
    id: 1,
    from: "Jakarta",
    to: "Bandung",
    IsPaidOff: false,
    users: [
      {
        id: 1,
        name: "asd"
      },
      {
        id: 2,
        name: "das"
      },
      {
        id: 3,
        name: "dasdad"
      }
    ]
  },
  flightData: [
    {
      go: {
        from: "13.00",
        to: "14.40",
        plane: "Lion",
        direct: "Direct"
      },
      return: {
        from: "13.00",
        to: "14.40",
        plane: "Lion",
        direct: "Direct"
      },
      price: "1.507.000"
    }
  ]
};
