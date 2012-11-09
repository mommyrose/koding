
CActivity = require '../../activity'
JCodeSnip = require './index'

module.exports = class CCodeSnipActivity extends CActivity

  @share()

  @set
    slugifyFrom     : 'title'
    encapsulatedBy  : CActivity
    sharedMethods   : CActivity.sharedMethods
    schema          : CActivity.schema
    relationships   :
      subject       :
        targetType  : JCodeSnip
        as          : 'content'