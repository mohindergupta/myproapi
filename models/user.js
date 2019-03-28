var mongoose = require('mongoose');
var Schema = mongoose.Schema;

var UserSchema = new Schema(
    {
      userid: Schema.Types.ObjectId,
      first_name: {type: String, required: true, max: 100},
      last_name: {type: String, required: true, max: 100},
      address1: {type: String, required: true, max: 100},
      address2: {type: String, required: false, max: 100},
      city: {type: String, required: true, max: 100},
      state: {type: String, required: true, max: 100},
      zipcode: {type: String, required: false, max: 100},
      country: {type: String, required: false, max: 100},
      email: {type: String, required: false, max: 100},
      phone: {type: String, required: false, max: 30},
      user_type: {type: String, required: false, max: 30, enum: ['Consumer','Service Provider']},
      date_of_birth: {type: Date, required: false},
      date_of_death: {type: Date, required: false},
      loginid: {type: String, required: false, max: 100},
      password: {type: String, required: false, max: 100}
    }
  );
  /* Commented by Mahender for now
  // Virtual for author's full name
  UserSchema
  .virtual('name')
  .get(function () {
    return this.family_name + ', ' + this.first_name;
  });
  
  // Virtual for author's lifespan
  AuthorSchema
  .virtual('lifespan')
  .get(function () {
    return (this.date_of_death.getYear() - this.date_of_birth.getYear()).toString();
  });
  
  // Virtual for author's URL
  AuthorSchema
  .virtual('url')
  .get(function () {
    return '/catalog/author/' + this._id;
  });
  */
  //Export model
  module.exports = mongoose.model('user', UserSchema);