<?php

namespace App\Admin\Models;

use Illuminate\Database\Eloquent\Model;
use App\Observers\ModelObserver;


class PoolAddress extends Model
{
    const TYPE_ADDRESS_COINBASE = 0;
    const TYPE_ADDRESS_PAYMENT = 1;
    const TYPE_ADDRESS_TRANSIT = 2;
    const TYPE_ADDRESS_OTHER = 3;

    public static $typeOptions = [
        self::TYPE_ADDRESS_COINBASE => "coinbase",
        self::TYPE_ADDRESS_PAYMENT => "支付",
        self::TYPE_ADDRESS_TRANSIT => "中转",
        self::TYPE_ADDRESS_OTHER => "其他"
    ];

    protected $table = 'pool_address';
    protected $fillable = ['address', 'coin_id', 'type'];

    public function pool()
    {
        $this->belongsTo(Pool::class, "pool_id");
    }

    public function coin()
    {
        $this->belongsTo(POWCoin::class, 'coin_id');
    }

    public static function boot()
    {
        parent::boot();
        $class = get_called_class();
        $class::observe(new ModelObserver());
    }
}
