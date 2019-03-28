var mongoose = require('mongoose');
var Schema = mongoose.Schema;

var BankSchema = new Schema(
    {
      objectype: {type: String, required: false, max: 100},
      bankid: Schema.Types.ObjectId,
      loginid: {type: String, required: false, max: 100},
      bank_profile_name: {type: String, required: false, max: 100},
      account_type: {type: String, required: false, max: 30, enum: ['','Checking','Saving','Money Market']},
      bank_name: {type: String, required: false, max: 100},
      routing_num: {type: String, required: false,max: 12},
      account_num: {type: String, required: false,max: 12},
      verification_result: {type: String, required: false, max: 100},
      shared_with_login: {type: String, required: false, max: 100}
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
  module.exports = mongoose.model('Bank', BankSchema);