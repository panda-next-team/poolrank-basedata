<?php

namespace App\Admin\Models;

use Illuminate\Database\Eloquent\Model;
use App\Observers\ModelObserver;


class PoolCoinbaseTag extends Model
{
    protected $table = 'pool_coinbase_tag';

    protected $fillable = ['tag'];

    public function pool() {
        $this->belongsTo(Pool::class,  "pool_id");
    }

    public static function boot()
    {
        parent::boot();
        $class = get_called_class();
        $class::observe(new ModelObserver());
    }
}
